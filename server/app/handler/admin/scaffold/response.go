package scaffold

import (
	"gf-admin/app/logic/system/scaffold"
	"gf-admin/app/model"
)

type (
	TableListRes []model.Table

	SingularRes string

	DatabaseListRes []model.Database

	ColumnListRes []model.Column

	PreviewRes []*scaffold.AutoCode
)
