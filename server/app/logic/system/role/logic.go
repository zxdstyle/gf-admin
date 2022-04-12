package role

import (
	"context"
	"gf-admin/app/logic/base"
	base2 "gf-admin/app/model/base"
	"gf-admin/app/model/system"
	"gf-admin/app/repository"
)

type Logic struct {
	*base.Logic
}

func NewLogic() *Logic {
	return &Logic{
		Logic: base.NewLogic(repository.Role()),
	}
}

func (Logic) Create(ctx context.Context, mo base2.RepositoryModel) error {
	if err := repository.Role().Create(ctx, mo); err != nil {
		return err
	}

	role := mo.(*system.Role)
	if role.Permissions != nil {
		return repository.Role().AttachPermissions(ctx, role, role.Permissions)
	}

	return nil
}

func (Logic) Update(ctx context.Context, mo base2.RepositoryModel) error {
	if err := repository.Role().Update(ctx, mo); err != nil {
		return err
	}

	role := mo.(*system.Role)
	if role.Permissions != nil {
		return repository.Role().SyncPermissions(ctx, role, role.Permissions)
	}

	return nil
}
