// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticlesDao is the data access object for table articles.
type ArticlesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ArticlesColumns // columns contains all the column names of Table for convenient usage.
}

// ArticlesColumns defines and stores column names for table articles.
type ArticlesColumns struct {
	Id          string //
	Title       string //
	CateId      string //
	Summary     string //
	OriginalUrl string //
	IsPublish   string //
	CreatedAt   string //
	UpdatedAt   string //
}

//  articlesColumns holds the columns for table articles.
var articlesColumns = ArticlesColumns{
	Id:          "id",
	Title:       "title",
	CateId:      "cate_id",
	Summary:     "summary",
	OriginalUrl: "original_url",
	IsPublish:   "is_publish",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewArticlesDao creates and returns a new DAO object for table data access.
func NewArticlesDao() *ArticlesDao {
	return &ArticlesDao{
		group:   "default",
		table:   "articles",
		columns: articlesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticlesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticlesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticlesDao) Columns() ArticlesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticlesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticlesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticlesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
