package system

import "gf-admin/app/model/base"

type (
	Permission struct {
		base.Model
		Name       *string `gorm:"column:name" json:"name,omitempty"`
		Slug       *string `gorm:"column:slug" json:"slug,omitempty"`
		HttpMethod *string `gorm:"column:http_method;default:''" json:"http_method,omitempty"`
		HttpPath   *string `gorm:"column:http_path;default:''" json:"http_path,omitempty"`
		Order      *uint   `gorm:"column:order;default:0" json:"order,omitempty"`
		ParentId   *uint   `gorm:"column:parent_id" json:"parent_id,omitempty"`

		Roles *Roles `gorm:"many2many:admin_role_permissions;" json:"roles,omitempty"`
	}

	Permissions []*Permission
)

func (Permission) TableName() string {
	return "admin_permissions"
}

func (*Permission) WithRules() base.PreloadRule {
	return base.PreloadRule{
		"Roles": base.DefaultPreloadFunc,
	}
}

func (p Permissions) Len() int {
	return len(p)
}

func (p Permissions) GetModel(i ...int) base.RepositoryModel {
	if len(p) == 0 {
		return &Permission{}
	}

	index := 0
	if len(i) > 0 {
		index = i[0]
	}
	return p[index]
}
