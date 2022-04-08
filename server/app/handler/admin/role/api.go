package role

import (
	"context"
	"gf-admin/app/logic"
	"gf-admin/app/model"
	"gf-admin/pkg/responses"
)

type Api struct {
}

func (Api) Index(ctx context.Context, req *IndexReq) (*responses.PaginatorRes, error) {
	paginator := responses.NewPaginator(&model.Roles{})
	err := logic.Role().Index(ctx, req.ListReq, paginator)
	return paginator, err
}

func (Api) Create(ctx context.Context, req *CreateReq) (responses.CreateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Role().Create(ctx, &mo)
	return mo, err
}

func (Api) Show(ctx context.Context, req *ShowReq) (responses.ShowRes, error) {
	var mo model.Role
	mo.SetID(req.ID)
	err := logic.Role().Show(ctx, req.GetWith(), &mo)
	return mo, err
}

func (Api) Update(ctx context.Context, req *UpdateReq) (responses.UpdateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Role().Update(ctx, &mo)
	return mo, err
}

func (Api) Destroy(ctx context.Context, req *DestroyReq) (responses.DestroyRes, error) {
	return nil, logic.Role().DestroyById(ctx, req.ID)
}
