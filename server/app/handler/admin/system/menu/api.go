package menu

import (
	"context"
	"gf-admin/app/logic"
	"gf-admin/app/model/system"
	"gf-admin/pkg/responses"
)

type Api struct {
}

func (Api) Index(ctx context.Context, req *IndexReq) (responses.IndexRes, error) {
	mos := &system.Menus{}
	if req.HasPage() {
		paginator := responses.NewPaginator(mos)
		err := logic.Menu().Index(ctx, req.ListReq, paginator)
		return paginator, err
	}

	err := logic.Menu().Tree(ctx, req.ListReq, mos)
	return mos, err
}

func (Api) Create(ctx context.Context, req *CreateReq) (responses.CreateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Menu().Create(ctx, &mo)
	return mo, err
}

func (Api) Show(ctx context.Context, req *ShowReq) (responses.ShowRes, error) {
	var mo system.Menu
	mo.SetID(req.ID)
	err := logic.Menu().Show(ctx, req.GetWith(), &mo)
	return mo, err
}

func (Api) Update(ctx context.Context, req *UpdateReq) (responses.UpdateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Menu().Update(ctx, &mo)
	return mo, err
}

func (Api) Destroy(ctx context.Context, req *DestroyReq) (responses.DestroyRes, error) {
	return nil, logic.Menu().DestroyById(ctx, req.ID)
}
