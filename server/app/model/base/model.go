package base

import (
	"github.com/gogf/gf/v2/os/gtime"
	"gorm.io/gorm"
)

var (
	DefaultPreloadFunc = func(tx *gorm.DB, resource string) *gorm.DB {
		return tx.Preload(resource)
	}
)

type (
	RepositoryModel interface {
		GetID() uint
		SetID(id uint)
	}

	RepositoryModels interface {
		Len() int
		GetModel(i ...int) RepositoryModel
	}

	HasWithModel interface {
		WithRules() PreloadRule
	}

	PreloadFunc func(tx *gorm.DB, resource string) *gorm.DB

	PreloadRule map[string]PreloadFunc

	Model struct {
		ID        uint        `json:"id"`
		CreatedAt *gtime.Time `json:"created_at"`
		UpdatedAt *gtime.Time `json:"updated_at"`
	}
)

func (m Model) GetID() uint {
	return m.ID
}

func (m *Model) SetID(id uint) {
	m.ID = id
}
