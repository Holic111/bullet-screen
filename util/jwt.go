package util

import (
	"bullet-screen/common"
	"bullet-screen/model/relative"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"time"
)

var control = &singleflight.Group{}

type JWT struct {
	SigningKey []byte
}


func NewJWT() *JWT {
	return &JWT{
		[]byte(common.Global_Jwt.JwtKey),
	}
}

// 创建token
func (j *JWT) CreateToken(claims relative.CustomClaims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 弃用旧token，发布新token
func (j *JWT) CreateTokenByOldToken(oldToken string,claims relative.CustomClaims) (string, error) {
	v, err, _ := control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*relative.CustomClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &relative.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				return nil, TOKEN_MALFORMED
			}else if ve.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, TOKEN_IS_EXPIRES
			}else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
				return nil, TOKEN_NOT_VALID
			}else {
				return nil, TOEKN_INVALID
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*relative.CustomClaims); ok && token.Valid {
			return claims, OK
		}
		return nil, TOKEN_NOT_VALID
	}
	return nil, TOEKN_INVALID
}

// 创建claims
func (j *JWT) CreateClaims(baseClaims relative.BaseClaims) relative.CustomClaims {
	claims := relative.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: common.Global_Jwt.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + common.Global_Jwt.ExpiresTime,
			Issuer: common.Global_Jwt.Issuer,
		},
	}
	return claims
}