package role

import (
	"context"
	"gf-admin/app/model/system"
	"gf-admin/app/repository/base"
)

type Repository interface {
	base.Repository
	AttachPermissions(ctx context.Context, role *system.Role, permissions *system.Permissions) error
	SyncPermissions(ctx context.Context, role *system.Role, permissions *system.Permissions) error
}
