package scaffold

import (
	"context"
	"gf-admin/app/model/system"
	"gf-admin/pkg"
	"gorm.io/gorm"
)

type DbRepository struct {
	Orm *gorm.DB
}

var _ Repository = &DbRepository{}

func NewDbRepository() *DbRepository {
	return &DbRepository{
		Orm: pkg.DB(),
	}
}

func (m DbRepository) GetTables(ctx context.Context, dbName string, data *[]system.Table) error {
	sql := "SELECT TABLE_NAME AS `name`, ENGINE AS `engine` FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = ?"
	return m.Orm.WithContext(ctx).Raw(sql, dbName).Find(data).Error
}

func (m DbRepository) GetColumns(ctx context.Context, tableName string, dbName string, data *[]system.Column) error {
	sql := `
	SELECT COLUMN_NAME AS name, DATA_TYPE AS type, COLUMN_COMMENT as comment, COLUMN_DEFAULT AS default_value,
		CASE IS_NULLABLE WHEN 'YES' THEN 1 ELSE 0 END AS nullable,
       	CASE DATA_TYPE
           WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'int' THEN c.NUMERIC_PRECISION
           WHEN 'bigint' THEN c.NUMERIC_PRECISION
           ELSE 0 END AS length
	FROM INFORMATION_SCHEMA.COLUMNS c
	WHERE table_name = ?
	  AND table_schema = ?
	`
	return m.Orm.WithContext(ctx).Raw(sql, tableName, dbName).Find(data).Error
}
