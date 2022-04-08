package guard

import "context"

type Guard interface {
	Token(ctx context.Context, identity uint) (string, error)
	ParseToken(ctx context.Context, token string) (uint, error)
}
