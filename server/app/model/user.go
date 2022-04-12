package model

import "gf-admin/app/model/base"

// User 用户表
type User struct {
	base.Model
	Email    *string `json:"email,omitempty"`                      //邮箱
	Username *string `json:"username,omitempty"`                   //账户
	Password *string `gorm:"default:''" json:"password,omitempty"` //密码
	IsActive *bool   `gorm:"default:0" json:"is_active,omitempty"` //是否启用
}

type Users []*User

func (*User) WithRules() base.PreloadRule {
	return base.PreloadRule{}
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) GetModel(i ...int) base.RepositoryModel {
	if len(u) == 0 {
		return &User{}
	}

	index := 0
	if len(i) > 0 {
		index = i[0]
	}
	return u[index]
}
