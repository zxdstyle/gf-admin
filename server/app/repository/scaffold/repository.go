package scaffold

import (
	"context"
	"gf-admin/app/model/system"
)

type Repository interface {
	GetTables(ctx context.Context, dbName string, data *[]system.Table) error
	GetColumns(ctx context.Context, tableName string, dbName string, data *[]system.Column) error
}
