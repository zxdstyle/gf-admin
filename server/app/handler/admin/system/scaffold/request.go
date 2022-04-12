package scaffold

import (
	"gf-admin/app/model/system"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	TableListReq struct {
		g.Meta `path:"/system/tables" method:"get"`
	}

	SingularReq struct {
		g.Meta `path:"/system/singular" method:"get"`
		Word   string `json:"word" v:"required"`
	}

	ColumnListReq struct {
		g.Meta `path:"/system/columns" method:"get"`
		Table  string `json:"table" v:"required"`
	}

	PreviewReq struct {
		g.Meta           `path:"/system/preview" method:"post"`
		Table            string        `json:"table"`
		Model            string        `json:"model"`
		Handler          string        `json:"handler"`
		ViewPath         string        `json:"view_path"`
		Repository       string        `json:"repository"`
		CreateModel      bool          `json:"create_model"`
		CreateRepository bool          `json:"create_repository"`
		CreateHandler    bool          `json:"create_handler"`
		CreateLang       bool          `json:"create_lang"`
		Fields           system.Fields `json:"fields"`
	}
)
