package permission

import (
	"gf-admin/app/logic/base"
	"gf-admin/app/repository"
)

type Logic struct {
	*base.Logic
}

func NewLogic() *Logic {
	return &Logic{
		Logic: base.NewLogic(repository.Permission()),
	}
}
