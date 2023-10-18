package jwt

import (
	"GyuBlog/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type BlogClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var blogSecret = []byte("GyuBlog")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return blogSecret, nil
}

const AccessTokenExpireDuration = time.Hour * 24
const RefreshTokenExpireDuration = time.Hour * 24 * 7

func GenToken(userID uint64, username string) (accessToken, refreshToken string, err error) {
	c := BlogClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			Issuer:    "GyuBlog",
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(blogSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
		Issuer:    "GyuBlog",
	}).SignedString(blogSecret)
	if err != nil {
		return "", "", err
	}

	return
}

func ParseToken(tokenString string) (claims *BlogClaims, err error) {
	var token *jwt.Token
	claims = new(BlogClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errcode.InvalidToken
	}
	return
}
