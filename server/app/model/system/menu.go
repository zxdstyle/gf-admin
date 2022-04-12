package system

import "gf-admin/app/model/base"

type (
	Menu struct {
		base.Model
		ParentId *uint   `gorm:"column:parent_id" json:"parent_id"`
		SortNum  *int    `gorm:"column:sort_num" json:"sort_num"`
		Title    *string `gorm:"column:title" json:"title"`
		Icon     *string `gorm:"column:icon" json:"icon"`
		Uri      *string `gorm:"column:uri" json:"uri"`

		Children *Menus `gorm:"-" json:"children,omitempty"`
	}

	Menus []*Menu
)

func (Menu) TableName() string {
	return "admin_menu"
}

func (m Menus) Len() int {
	return len(m)
}

func (m Menus) GetModel(i ...int) base.RepositoryModel {
	if len(m) == 0 {
		return &Menu{}
	}

	index := 0
	if len(i) > 0 {
		index = i[0]
	}
	return m[index]
}
