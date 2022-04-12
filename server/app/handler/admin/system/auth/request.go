package auth

import "github.com/gogf/gf/v2/frame/g"

type (
	LoginReq struct {
		g.Meta   `path:"/login" method:"post"`
		Email    string `json:"email" v:"required|email"`
		Password string `json:"password" v:"required"`
	}
)
