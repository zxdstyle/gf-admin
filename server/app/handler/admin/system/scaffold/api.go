package scaffold

import (
	"context"
	"gf-admin/app/logic"
	"gf-admin/app/model/system"
	"gf-admin/app/types"
	"gf-admin/pkg/responses"
	"gf-admin/pkg/utils"
)

type Scaffold struct {
}

func (Scaffold) Table(ctx context.Context, req *TableListReq) (*types.OptionsRes, error) {
	var t types.Options
	if err := logic.Scaffold().GetTables(ctx, &t); err != nil {
		return nil, responses.Failed(err)
	}
	res := types.OptionsRes(t)
	return &res, nil
}

func (Scaffold) Singular(ctx context.Context, req *SingularReq) (*SingularRes, error) {
	res := SingularRes(utils.Singular(req.Word))
	return &res, nil
}

func (Scaffold) Column(ctx context.Context, req *ColumnListReq) (*ColumnListRes, error) {
	var t []system.Column
	if err := logic.Scaffold().GetColumns(ctx, req.Table, &t); err != nil {
		return nil, responses.Failed(err)
	}
	res := ColumnListRes(t)
	return &res, nil
}

func (Scaffold) Preview(ctx context.Context, req *PreviewReq) (*PreviewRes, error) {
	return nil, logic.Scaffold().CreateCode(ctx, req.ViewPath, req.Table, req.Fields)
}
