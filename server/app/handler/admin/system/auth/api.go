package auth

import (
	"context"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/app/repository"
	"gf-admin/app/service"
	"gf-admin/app/utils"
	"gorm.io/gorm"
)

type Auth struct {
}

func (Auth) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	user := &model.User{
		Email: &req.Email,
	}
	if err := repository.User().First(ctx, user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("账号或密码错误")
		}
		return nil, err
	}
	if !utils.PasswordVerify(req.Password, *user.Password) {
		return nil, fmt.Errorf("账号或密码错误")
	}
	token, err := service.Guard.Token(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &LoginRes{Token: token}, nil
}
