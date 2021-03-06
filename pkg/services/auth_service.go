package services

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/berkayersoyy/go-products-example/pkg/utils/config"
	jwtutils "github.com/berkayersoyy/go-products-example/pkg/utils/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"
)

type authService struct {
	Client *redis.Client
}

type AuthService interface {
	ValidateToken(r *http.Request) (*jwt.Token, error)
	TokenValid(r *http.Request) error
	CreateToken(userid uint) (*jwtutils.TokenDetails, error)
	CreateAuth(userid uint, td *jwtutils.TokenDetails) error
	DeleteTokens(authD *jwtutils.AccessDetails) error
	ExtractTokenMetadata(r *http.Request) (*jwtutils.AccessDetails, error)
	FetchAuth(authD *jwtutils.AccessDetails) (uint64, error)
	DeleteAuth(givenUuid string) (int64, error)
}

func ProvideAuthService(c *redis.Client) AuthService {
	return &authService{Client: c}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (a *authService) ValidateToken(r *http.Request) (*jwt.Token, error) {
	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (a *authService) TokenValid(r *http.Request) error {
	token, err := a.ValidateToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}
func (a *authService) CreateToken(userid uint) (*jwtutils.TokenDetails, error) {
	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	td := &jwtutils.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(userid))

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(conf.AccessSecret))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(conf.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (a *authService) CreateAuth(userid uint, td *jwtutils.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := a.Client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := a.Client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
func (a *authService) ExtractTokenMetadata(r *http.Request) (*jwtutils.AccessDetails, error) {
	token, err := a.ValidateToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.Atoi(fmt.Sprintf("%.f", claims["user_id"]))
		if err != nil {
			return nil, err
		}
		return &jwtutils.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}
func (a *authService) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := a.Client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
func (a *authService) DeleteTokens(authD *jwtutils.AccessDetails) error {
	refreshUuid := fmt.Sprintf("%s++%d", authD.AccessUuid, authD.UserId)
	deletedAt, err := a.Client.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := a.Client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
func (a *authService) FetchAuth(authD *jwtutils.AccessDetails) (uint64, error) {
	userid, err := a.Client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	if uint64(authD.UserId) != userID {
		return 0, errors.New("unauthorized")
	}
	return userID, nil
}
