package model

type (
	Role struct {
		Model
		Name *string `gorm:"column:name" json:"name,omitempty"`
		Slug *string `gorm:"column:slug" json:"slug,omitempty"`

		Permissions *Permissions `gorm:"many2many:admin_role_permissions;" json:"permissions,omitempty"`
	}

	Roles []*Role
)

func (Role) TableName() string {
	return "admin_roles"
}

func (*Role) WithRules() PreloadRule {
	return PreloadRule{
		"Permissions": defaultPreloadFunc,
	}
}

func (r Roles) Len() int {
	return len(r)
}

func (r Roles) GetModel(i ...int) RepositoryModel {
	if len(r) == 0 {
		return &Role{}
	}

	index := 0
	if len(i) > 0 {
		index = i[0]
	}
	return r[index]
}
