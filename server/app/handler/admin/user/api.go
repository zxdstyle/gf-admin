package user

import (
	"context"
	"gf-admin/app/logic"
	"gf-admin/app/model"
	"gf-admin/pkg/responses"
)

type Api struct {
}

func (Api) Index(ctx context.Context, req *IndexReq) (*responses.PaginatorRes, error) {
	paginator := responses.NewPaginator(&model.Users{})
	err := logic.User().Index(ctx, req.ListReq, paginator)
	return paginator, err
}

func (Api) Create(ctx context.Context, req *CreateReq) (responses.CreateRes, error) {
	user, err := req.ToModel()
	if err != nil {
		return nil, err
	}
	err = logic.User().Create(ctx, &user)
	return user, err
}

func (Api) Show(ctx context.Context, req *ShowReq) (responses.ShowRes, error) {
	var user model.User
	user.SetID(req.ID)
	err := logic.User().Show(ctx, req.GetWith(), &user)
	return user, err
}

func (Api) Update(ctx context.Context, req *UpdateReq) (responses.UpdateRes, error) {
	user, err := req.ToModel()
	if err != nil {
		return nil, err
	}
	err = logic.User().Update(ctx, &user)
	return user, err
}

func (Api) Destroy(ctx context.Context, req *DestroyReq) (responses.DestroyRes, error) {
	return nil, logic.User().DestroyById(ctx, req.ID)
}
