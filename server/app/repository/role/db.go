package role

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/repository/base"
	"gf-admin/pkg"
)

type DbRepository struct {
	*base.GormRepository
}

func NewDbRepository() *DbRepository {
	return &DbRepository{
		GormRepository: base.NewGormRepository(pkg.DB().Model(model.Role{})),
	}
}

// AttachPermissions 添加权限
func (repo *DbRepository) AttachPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error {
	if role == nil || permissions == nil {
		return nil
	}

	if err := repo.Orm.WithContext(ctx).Model(role).Omit("Permissions.*").Association("Permissions").Append(permissions); err != nil {
		return err
	}

	if err := repo.Show(ctx, []string{"Permissions"}, role); err != nil {
		return err
	}

	return nil
}

// SyncPermissions 同步权限
func (repo *DbRepository) SyncPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error {
	if role == nil || permissions == nil {
		return nil
	}

	tx := repo.Orm.WithContext(ctx).Begin()

	if err := tx.Model(role).Association("Permissions").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(role).Omit("Permissions.*").Association("Permissions").Append(permissions); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	if err := repo.Show(ctx, []string{"Permissions"}, role); err != nil {
		return err
	}

	return nil
}
