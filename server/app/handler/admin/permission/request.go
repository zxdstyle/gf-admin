package permission

import (
	"gf-admin/app/model"
	"gf-admin/pkg/request"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IndexReq struct {
		g.Meta `path:"/permissions" method:"get"`
		request.ListReq
	}

	ShowReq struct {
		g.Meta `path:"/permissions/:id" method:"get"`
		ID     uint `json:"id"`
		request.ShowReq
	}

	CreateReq struct {
		g.Meta     `path:"/permissions" method:"post"`
		Name       *string `json:"name,omitempty" v:"required"`
		Slug       *string `json:"slug,omitempty" v:"required|unique:admin_permissions,slug#缺少权限标识|该权限标识已被使用"`
		HttpMethod *string `json:"http_method,omitempty"`
		HttpPath   *string `json:"http_path,omitempty"`
		Order      *uint   `json:"order,omitempty"`
		ParentId   *uint   `json:"parent_id,omitempty" v:"required|exists:admin_permissions,id#缺少父级权限ID|父级权限ID不存在"`
	}

	UpdateReq struct {
		g.Meta     `path:"/permissions/:id" method:"put"`
		ID         uint    `json:"id" v:"required"`
		Name       *string `json:"name,omitempty"`
		Slug       *string `json:"slug,omitempty" v:"unique:admin_permissions,slug#该权限标识已被使用"`
		HttpMethod *string `json:"http_method,omitempty"`
		HttpPath   *string `json:"http_path,omitempty"`
		Order      *uint   `json:"order,omitempty"`
		ParentId   *uint   `json:"parent_id,omitempty" v:"exists:admin_permissions,id#父级权限ID不存在"`
	}

	DestroyReq struct {
		g.Meta `path:"/permissions/:id" method:"delete"`
		ID     uint `json:"id" v:"required"`
	}
)

func (req CreateReq) ToModel() (model.Permission, error) {
	return model.Permission{
		Name:       req.Name,
		Slug:       req.Slug,
		HttpMethod: req.HttpMethod,
		HttpPath:   req.HttpPath,
		Order:      req.Order,
		ParentId:   req.ParentId,
	}, nil
}

func (req UpdateReq) ToModel() (model.Permission, error) {
	role := model.Permission{
		Name:       req.Name,
		Slug:       req.Slug,
		HttpMethod: req.HttpMethod,
		HttpPath:   req.HttpPath,
		Order:      req.Order,
		ParentId:   req.ParentId,
	}
	role.SetID(req.ID)
	return role, nil
}
