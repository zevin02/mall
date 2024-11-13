package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//jwt是一种开放，用于网络应用环境中，安全传递信息，主要用于身份验证

var jwtSecret = []byte("jwt_secret") //签名token

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) //must be expired or be dangerous
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //设置过期时间
			Issuer:    "FanOne-Mall",
		},
	}
	//创建一个jwt token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //token生成
	return token, err

}
