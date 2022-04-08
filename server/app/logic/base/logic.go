package base

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/repository/base"
	"gf-admin/pkg/request"
	"gf-admin/pkg/responses"
)

type Logic struct {
	repo base.Repository
}

func NewLogic(repo base.Repository) *Logic {
	return &Logic{repo: repo}
}

func (l Logic) Index(ctx context.Context, req request.ListReq, paginator *responses.PaginatorRes) error {
	return l.repo.Index(ctx, req, paginator)
}

func (l Logic) Show(ctx context.Context, with []string, mo model.RepositoryModel) error {
	return l.repo.Show(ctx, with, mo)
}

func (l Logic) Create(ctx context.Context, mo model.RepositoryModel) error {
	return l.repo.Create(ctx, mo)
}

func (l Logic) Update(ctx context.Context, mo model.RepositoryModel) error {
	return l.repo.Update(ctx, mo)
}

func (l Logic) Destroy(ctx context.Context, mo model.RepositoryModel) error {
	return l.repo.Destroy(ctx, mo)
}

func (l Logic) DestroyById(ctx context.Context, id uint) error {
	return l.repo.DestroyById(ctx, id)
}
