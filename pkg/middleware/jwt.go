package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("daitoue_secret_key")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func ReleaseToken(id int64) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //终止时间，7天后
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wuhlan3.tech",
			Subject:   "user token",
		},
	}
	// 这里需要注意当使用jwt.SigningMethodHS256方式生成token串时
	// SignedString方法的参数应该是[]byte数组
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
