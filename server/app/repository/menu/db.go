package menu

import (
	"gf-admin/app/model/system"
	"gf-admin/app/repository/base"
	"gf-admin/pkg"
)

type DbRepository struct {
	*base.GormRepository
}

func NewDbRepository() *DbRepository {
	return &DbRepository{
		GormRepository: base.NewGormRepository(pkg.DB().Model(system.Menu{})),
	}
}
