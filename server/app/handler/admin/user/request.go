package user

import (
	"gf-admin/app/model"
	"gf-admin/app/utils"
	"gf-admin/pkg/request"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IndexReq struct {
		g.Meta `path:"/users" method:"get"`
		request.ListReq
	}

	ShowReq struct {
		g.Meta `path:"/users/:id" method:"get"`
		ID     uint `json:"id"`
		request.ShowReq
	}

	CreateReq struct {
		g.Meta   `path:"/users" method:"post"`
		Email    *string `json:"email" v:"required|email"` //邮箱
		Username *string `json:"username" v:"required"`    //账户
		Password *string `json:"password"`                 //密码
		IsActive *bool   `json:"is_active"`                //是否启用
	}

	UpdateReq struct {
		g.Meta   `path:"/users/:id" method:"put"`
		ID       uint    `json:"id" v:"required"`
		Email    *string `json:"email" v:"email"` //邮箱
		Username *string `json:"username"`        //账户
		Password *string `json:"password"`        //密码
		IsActive *bool   `json:"is_active"`       //是否启用
	}

	DestroyReq struct {
		g.Meta `path:"/users/:id" method:"delete"`
		ID     uint `json:"id" v:"required"`
	}
)

func (req CreateReq) ToModel() (model.User, error) {
	active := new(bool)
	if req.IsActive != nil {
		active = req.IsActive
	}
	var (
		pwd string
		err error
	)
	if req.Password != nil {
		pwd, err = utils.PasswordHash(*req.Password)
		if err != nil {
			return model.User{}, err
		}
	}

	return model.User{
		Email:    req.Email,
		Username: req.Username,
		Password: &pwd,
		IsActive: active,
	}, nil
}

func (req UpdateReq) ToModel() (model.User, error) {
	var (
		pwd string
		err error
	)
	if req.Password != nil {
		pwd, err = utils.PasswordHash(*req.Password)
		if err != nil {
			return model.User{}, err
		}
	}
	user := model.User{
		Email:    req.Email,
		Username: req.Username,
		Password: &pwd,
		IsActive: req.IsActive,
	}
	user.SetID(req.ID)
	return user, nil
}
