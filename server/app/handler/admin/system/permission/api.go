package permission

import (
	"context"
	"gf-admin/app/logic"
	"gf-admin/app/model/system"
	"gf-admin/pkg/responses"
)

type Api struct {
}

func (Api) Index(ctx context.Context, req *IndexReq) (responses.IndexRes, error) {
	mos := &system.Permissions{}
	if req.HasPage() {
		paginator := responses.NewPaginator(mos)
		err := logic.Permission().Index(ctx, req.ListReq, paginator)
		return paginator, err
	}

	err := logic.Permission().IndexAll(ctx, req.ListReq, mos)
	return mos, err
}

func (Api) Create(ctx context.Context, req *CreateReq) (responses.CreateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Permission().Create(ctx, &mo)
	return mo, err
}

func (Api) Show(ctx context.Context, req *ShowReq) (responses.ShowRes, error) {
	var mo system.Permission
	mo.SetID(req.ID)

	err := logic.Permission().Show(ctx, req.GetWith(), &mo)
	return mo, err
}

func (Api) Update(ctx context.Context, req *UpdateReq) (responses.UpdateRes, error) {
	mo, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	err = logic.Permission().Update(ctx, &mo)
	return mo, err
}

func (Api) Destroy(ctx context.Context, req *DestroyReq) (responses.DestroyRes, error) {
	return nil, logic.Permission().DestroyById(ctx, req.ID)
}
