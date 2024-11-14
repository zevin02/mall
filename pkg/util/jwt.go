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

// 签发token
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

// ParseToken 验证用户的token，解析JWT令牌并返回声明的信息
func ParseToken(token string) (*Claims, error) {
	//
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//检查解析是否成功
	if tokenClaims != nil {
		//尝试将Claims转化成Claims类型，并检查令牌的有小型
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//解析失败
	return nil, err
}
