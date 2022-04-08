package base

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/pkg/request"
	"gf-admin/pkg/responses"
)

type Repository interface {
	Index(ctx context.Context, req request.ListReq, paginator *responses.PaginatorRes) error
	Show(ctx context.Context, with []string, mo model.RepositoryModel) error
	Create(ctx context.Context, mo model.RepositoryModel) error
	Update(ctx context.Context, mo model.RepositoryModel) error
	Destroy(ctx context.Context, mo model.RepositoryModel) error
	DestroyById(ctx context.Context, id uint) error
	First(ctx context.Context, mo model.RepositoryModel) error
}
