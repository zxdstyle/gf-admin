package utils

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
	singleFlight     = &singleflight.Group{}
)

type (
	JWT struct {
		SigningKey []byte
	}

	CustomClaims struct {
		jwt.RegisteredClaims
		Identity   uint
		BufferTime int64
	}
)

func NewJWT() *JWT {
	return &JWT{
		g.Cfg().MustGet(context.Background(), "jwt.key", "gf-admin").Bytes(),
	}
}

func (j *JWT) CreateClaims(ctx context.Context, identity uint) CustomClaims {
	ttl := g.Cfg().MustGet(ctx, "jwt.ttl").Duration()
	claims := CustomClaims{
		Identity:   identity,
		BufferTime: g.Cfg().MustGet(ctx, "jwt.ttl").Int64(), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-10 * time.Second)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),               // 过期时间
			Issuer:    g.Cfg().MustGet(ctx, "app_name", "gf-admin").String(), // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := singleFlight.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
