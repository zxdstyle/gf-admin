package request

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"gorm.io/gorm"
)

const (
	defaultPageSize = 20
)

type (
	ListReq struct {
		Page     int      `json:"page"`
		PageSize int      `json:"pageSize"`
		With     []string `json:"with"`
		Selects  string   `json:"selects"`
		OrderBy  string   `json:"order_by"`
		Order    string   `json:"order" v:"in:asc,desc,ASC,DESC"`
	}

	ShowReq struct {
		_with   []string
		With    string `json:"with"`
		Selects string `json:"selects"`
	}
)

func (req *ListReq) ToListQuery(query *gorm.DB) *gorm.DB {
	if len(req.Selects) > 0 {
		query = query.Select(req.Selects)
	}

	sort := "DESC"
	if len(req.Order) > 0 {
		sort = req.Order
	}
	orderBy := "id"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	order := fmt.Sprintf("%s %s", orderBy, sort)

	query = query.Order(order).Limit(req.GetPageSize()).Offset(req.GetOffset())

	return query
}

func (req *ListReq) ToCountQuery(ctx context.Context, query *gorm.DB, count *int64) error {
	return req.queryByFilter(query).Count(count).Error
}

func (req *ListReq) GetPageSize() int {
	if req.PageSize == 0 {
		return defaultPageSize
	}
	return req.PageSize
}

func (req *ListReq) GetPage() int {
	if req.Page == 0 {
		return 1
	}
	return req.Page
}

func (req *ListReq) GetOffset() int {
	return (req.GetPage() - 1) * req.GetPageSize()
}

func (req *ListReq) queryByFilter(tx *gorm.DB) *gorm.DB {
	// TODO
	return tx
}

func (req *ShowReq) GetWith() []string {
	if len(req.With) == 0 {
		return []string{}
	}
	if len(req._with) == 0 {
		req._with = gstr.Split(req.With, ",")
	}
	return req._with
}
