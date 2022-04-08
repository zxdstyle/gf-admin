package user

import (
	"gf-admin/app/model"
	"gf-admin/app/repository/base"
	"gf-admin/pkg"
)

type DbRepository struct {
	*base.GormRepository
}

func NewDbRepository() *DbRepository {
	return &DbRepository{
		GormRepository: base.NewGormRepository(pkg.DB().Model(model.User{})),
	}
}
