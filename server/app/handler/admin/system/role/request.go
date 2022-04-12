package role

import (
	"gf-admin/app/model/system"
	"gf-admin/pkg/request"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IndexReq struct {
		g.Meta `path:"/roles" method:"get"`
		request.ListReq
	}

	ShowReq struct {
		g.Meta `path:"/roles/:id" method:"get"`
		request.ShowReq
		ID uint `json:"id"`
	}

	CreateReq struct {
		g.Meta      `path:"/roles" method:"post"`
		Name        *string             `json:"email" v:"required"`
		Slug        *string             `json:"username" v:"required|unique:admin_roles,slug#缺少角色标识|角色标识已被使用"`
		Permissions *system.Permissions `json:"permissions"`
	}

	UpdateReq struct {
		g.Meta      `path:"/roles/:id" method:"put"`
		ID          uint                `json:"id" v:"required"`
		Name        *string             `json:"email"`
		Slug        *string             `json:"username" v:"unique:admin_roles,slug#角色标识已被使用"`
		Permissions *system.Permissions `json:"permissions"`
	}

	DestroyReq struct {
		g.Meta `path:"/roles/:id" method:"delete"`
		ID     uint `json:"id" v:"required"`
	}
)

func (req CreateReq) ToModel() (system.Role, error) {
	return system.Role{
		Name:        req.Name,
		Slug:        req.Slug,
		Permissions: req.Permissions,
	}, nil
}

func (req UpdateReq) ToModel() (system.Role, error) {
	role := system.Role{
		Name:        req.Name,
		Slug:        req.Slug,
		Permissions: req.Permissions,
	}
	role.SetID(req.ID)
	return role, nil
}
