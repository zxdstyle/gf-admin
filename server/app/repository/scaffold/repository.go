package scaffold

import (
	"context"
	"gf-admin/app/model"
)

type Repository interface {
	GetTables(ctx context.Context, dbName string, data *[]model.Table) error
	GetColumns(ctx context.Context, tableName string, dbName string, data *[]model.Column) error
}
