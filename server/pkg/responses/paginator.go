package responses

const (
	DefaultPageSize = 20
)

type (
	PaginatorRes struct {
		Page     int         `json:"page"`
		PageSize int         `json:"pageSize"`
		Data     interface{} `json:"data"`
		Total    int64       `json:"total"`
	}
)

func NewPaginator(data interface{}) *PaginatorRes {
	return &PaginatorRes{
		Data: data,
		Page: 1,
	}
}
