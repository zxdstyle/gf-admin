package base

import (
	"context"
	"gf-admin/app/model/base"
	"gf-admin/pkg/request"
	"gf-admin/pkg/responses"
	"github.com/gogf/gf/v2/text/gstr"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository struct {
	Orm *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		Orm: db,
	}
}

func (repo GormRepository) Index(ctx context.Context, req request.ListReq, paginator *responses.PaginatorRes) error {
	tx := repo.Orm.WithContext(ctx)

	if err := req.ToCountQuery(ctx, tx, &paginator.Total); err != nil {
		return err
	}

	if len(req.With) > 0 {
		if val, ok := paginator.Data.(base.RepositoryModels); ok {
			mo := val.GetModel()
			tx = repo.preloadResource(tx, mo, req.GetWith())
		}
	}
	if err := req.ToListQuery(tx).Find(paginator.Data).Error; err != nil {
		return err
	}
	paginator.PageSize = req.GetPageSize()
	paginator.Page = req.GetPage()
	return nil
}

func (repo GormRepository) IndexAll(ctx context.Context, req request.ListReq, mos base.RepositoryModels) error {
	tx := repo.Orm.WithContext(ctx)

	if len(req.With) > 0 {
		mo := mos.GetModel()
		tx = repo.preloadResource(tx, mo, req.GetWith())
	}

	return req.ToListQuery(tx).Find(mos).Error
}

func (repo GormRepository) Show(ctx context.Context, with []string, mo base.RepositoryModel) error {
	tx := repo.Orm.WithContext(ctx)

	tx = repo.preloadResource(tx, mo, with)

	return tx.First(mo).Error
}

func (repo GormRepository) Create(ctx context.Context, mo base.RepositoryModel) error {
	if err := repo.Orm.WithContext(ctx).Omit(clause.Associations).Create(mo).Error; err != nil {
		return err
	}
	return repo.Show(ctx, []string{}, mo)
}

func (repo GormRepository) Update(ctx context.Context, mo base.RepositoryModel) error {
	if err := repo.Orm.WithContext(ctx).Where("id = ?", mo.GetID()).Updates(mo).Error; err != nil {
		return err
	}
	return repo.Show(ctx, []string{}, mo)
}

func (repo GormRepository) Destroy(ctx context.Context, mo base.RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Delete(mo).Error
}

func (repo GormRepository) DestroyById(ctx context.Context, id uint) error {
	return repo.Orm.WithContext(ctx).Where("id = ?", id).Delete(&repo.Orm.Statement.Model).Error
}

func (repo GormRepository) First(ctx context.Context, mo base.RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Where(mo).First(mo).Error
}

func (repo GormRepository) preloadResource(tx *gorm.DB, mo base.RepositoryModel, with []string) *gorm.DB {
	if val, ok := mo.(base.HasWithModel); ok {
		rules := val.WithRules()
		for _, res := range with {
			res = gstr.UcFirst(res)
			if preload, o := rules[res]; o {
				tx = preload(tx, res)
			}
		}
	}
	return tx
}
