package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secretKey")

type Claims struct {
	UserName string
	jwt.StandardClaims
}

//生成JWT
func ReleaseToken(userName string) (string, error) {
	expirationTime := time.Now().Add(1 * 24 * time.Hour) //定义过期时间
	claims := &Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin.practise",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//tokenString中截取token并返回，解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
