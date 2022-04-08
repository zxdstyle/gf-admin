package role

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/repository/base"
)

type Repository interface {
	base.Repository
	AttachPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
	SyncPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
}
