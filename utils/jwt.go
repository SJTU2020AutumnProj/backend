package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("wearesjtuboxin2020")

/*
Claims : payload of jwt
*/
type Claims struct {
	UserID   int32  `json:"userid"`
	UserName string `json:"username"`
	Password string `json:"password"`
	UserType int32  `json:"usertype"`
	jwt.StandardClaims
}

/*
JWTSign : sign a JWT token
*/
func JWTSign(userID int32, userName string, password string, userType int32) (string, error) {
	expireMinute := 60
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expireMinute) * time.Minute)
	claims := Claims{
		UserID:   userID,
		UserName: userName,
		Password: password,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "boxin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

/*
JWTVerify : parse JWT token and return payload information
*/
func JWTVerify(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil && tokenClaims.Valid {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
