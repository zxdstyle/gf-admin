package base

import (
	"context"
	"gf-admin/app/model/base"
	"gf-admin/pkg/request"
	"gf-admin/pkg/responses"
)

type Repository interface {
	Index(ctx context.Context, req request.ListReq, paginator *responses.PaginatorRes) error
	IndexAll(ctx context.Context, req request.ListReq, mos base.RepositoryModels) error
	Show(ctx context.Context, with []string, mo base.RepositoryModel) error
	Create(ctx context.Context, mo base.RepositoryModel) error
	Update(ctx context.Context, mo base.RepositoryModel) error
	Destroy(ctx context.Context, mo base.RepositoryModel) error
	DestroyById(ctx context.Context, id uint) error
	First(ctx context.Context, mo base.RepositoryModel) error
}
