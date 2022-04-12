package menu

import (
	"gf-admin/app/model/system"
	"gf-admin/pkg/request"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IndexReq struct {
		g.Meta `path:"/menus" method:"get"`
		request.ListReq
	}

	ShowReq struct {
		g.Meta `path:"/menus/:id" method:"get"`
		ID     uint `json:"id"`
		request.ShowReq
	}

	CreateReq struct {
		g.Meta   `path:"/menus" method:"post"`
		ParentId *uint   `json:"parent_id" v:"required|exists:admin_menu,id,0#请选择父级菜单|父级菜单不存在"`
		SortNum  *int    `json:"sort_num" v:"min:0"`
		Title    *string `json:"title" v:"required"`
		Icon     *string `json:"icon" v:"required"`
		Uri      *string `json:"uri"`
	}

	UpdateReq struct {
		g.Meta   `path:"/menus/:id" method:"put"`
		ID       uint    `json:"id" v:"required"`
		ParentId *uint   `json:"parent_id" v:"exists:admin_menu,id,0#父级菜单不存在"`
		SortNum  *int    `json:"sort_num" v:"min:0"`
		Title    *string `json:"title"`
		Icon     *string `json:"icon"`
		Uri      *string `json:"uri"`
	}

	DestroyReq struct {
		g.Meta `path:"/menus/:id" method:"delete"`
		ID     uint `json:"id" v:"required"`
	}
)

func (req CreateReq) ToModel() (system.Menu, error) {
	return system.Menu{
		ParentId: req.ParentId,
		SortNum:  req.SortNum,
		Title:    req.Title,
		Icon:     req.Icon,
		Uri:      req.Uri,
	}, nil
}

func (req UpdateReq) ToModel() (system.Menu, error) {
	role := system.Menu{
		ParentId: req.ParentId,
		SortNum:  req.SortNum,
		Title:    req.Title,
		Icon:     req.Icon,
		Uri:      req.Uri,
	}
	role.SetID(req.ID)
	return role, nil
}
