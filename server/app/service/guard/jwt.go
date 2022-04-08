package guard

import (
	"context"
	"gf-admin/app/utils"
)

type JwtGuard struct {
	jwt *utils.JWT
}

func NewJwtGuard() *JwtGuard {
	return &JwtGuard{
		jwt: utils.NewJWT(),
	}
}

func (guard *JwtGuard) Token(ctx context.Context, identity uint) (string, error) {
	claims := guard.jwt.CreateClaims(ctx, identity)
	return guard.jwt.CreateToken(claims)
}

func (guard *JwtGuard) ParseToken(ctx context.Context, token string) (uint, error) {
	claims, err := guard.jwt.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return claims.Identity, nil
}
