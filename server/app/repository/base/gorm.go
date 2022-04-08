package base

import (
	"context"
	"gf-admin/app/model"
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
	query := repo.Orm.WithContext(ctx)

	if err := req.ToCountQuery(ctx, query, &paginator.Total); err != nil {
		return err
	}

	if len(req.With) > 0 {
		if val, ok := paginator.Data.(model.RepositoryModels); ok {
			mo := val.GetModel()
			if m, o := mo.(model.HasWithModel); o {
				rules := m.WithRules()
				for _, res := range req.With {
					if preload, k := rules[res]; k {
						query = preload(query, res)
					}
				}
			}
		}
	}
	if err := req.ToListQuery(query).Find(paginator.Data).Error; err != nil {
		return err
	}
	paginator.PageSize = req.GetPageSize()
	paginator.Page = req.GetPage()
	return nil
}

func (repo GormRepository) Show(ctx context.Context, with []string, mo model.RepositoryModel) error {
	tx := repo.Orm.WithContext(ctx)

	if val, ok := mo.(model.HasWithModel); ok {
		rules := val.WithRules()
		for _, res := range with {
			res = gstr.UcFirst(res)
			if preload, o := rules[res]; o {
				tx = preload(tx, res)
			}
		}
	}

	return tx.First(mo).Error
}

func (repo GormRepository) Create(ctx context.Context, mo model.RepositoryModel) error {
	if err := repo.Orm.WithContext(ctx).Omit(clause.Associations).Create(mo).Error; err != nil {
		return err
	}
	return repo.Show(ctx, []string{}, mo)
}

func (repo GormRepository) Update(ctx context.Context, mo model.RepositoryModel) error {
	if err := repo.Orm.WithContext(ctx).Where("id = ?", mo.GetID()).Updates(mo).Error; err != nil {
		return err
	}
	return repo.Show(ctx, []string{}, mo)
}

func (repo GormRepository) Destroy(ctx context.Context, mo model.RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Delete(mo).Error
}

func (repo GormRepository) DestroyById(ctx context.Context, id uint) error {
	return repo.Orm.WithContext(ctx).Where("id = ?", id).Delete(&repo.Orm.Statement.Model).Error
}

func (repo GormRepository) First(ctx context.Context, mo model.RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Where(mo).First(mo).Error
}
