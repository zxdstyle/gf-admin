package scaffold

import (
	"context"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/app/repository"
	"gf-admin/app/types"
	"gf-admin/pkg/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"sync"
)

const (
	apiVersion  = "v1"
	apiPageTpl  = "stubs/web/api.ts.tpl"
	useDataTpl  = "stubs/web/useData.tsx.tpl"
	listPageTpl = "stubs/web/list.vue.tpl"
	formPageTpl = "stubs/web/form.vue.tpl"
)

type (
	Logic struct {
		viewer *gview.View
	}

	AutoCode struct {
		Lang     string     `json:"lang"`
		Path     string     `json:"path"`
		FullPath string     `json:"-"`
		Content  string     `json:"content,omitempty"`
		Children []AutoCode `json:"children,omitempty"`
	}
)

func NewLogic() *Logic {
	viewer := gview.New()
	if err := viewer.SetConfig(gview.Config{
		Delimiters: []string{"${", "}"},
	}); err != nil {
		g.Log().Error(context.TODO(), err)
	}
	return &Logic{
		viewer: viewer,
	}
}

func (Logic) GetTables(ctx context.Context, options *types.Options) error {
	var tables []model.Table
	if err := repository.Scaffold().GetTables(ctx, g.DB().GetConfig().Name, &tables); err != nil {
		return err
	}

	for _, table := range tables {
		*options = append(*options, types.Option{
			Label: table.Name,
			Value: table.Name,
		})
	}
	return nil
}

func (Logic) GetColumns(ctx context.Context, table string, columns *[]model.Column) error {
	if err := repository.Scaffold().GetColumns(ctx, table, g.DB().GetConfig().Name, columns); err != nil {
		return err
	}
	for i, column := range *columns {
		if column.Comment == "" {
			(*columns)[i].Comment = gstr.UcFirst(column.Name)
		}
	}
	return nil
}

func (l Logic) CreateCode(ctx context.Context, path, table string, fields model.Fields) error {
	wg := &sync.WaitGroup{}

	var errors []error

	wg.Add(1)
	go func() {
		if err := l.CreateListPage(ctx, fields, table, path); err != nil {
			errors = append(errors, err)
			g.Log().Error(ctx, err)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := l.CreateApiPage(ctx, table); err != nil {
			errors = append(errors, err)
			g.Log().Error(ctx, err)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := l.CreateFormPage(ctx, fields, table, path); err != nil {
			errors = append(errors, err)
			g.Log().Error(ctx, err)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := l.CreateDataPage(ctx, fields, table, path); err != nil {
			errors = append(errors, err)
			g.Log().Error(ctx, err)
		}
		wg.Done()
	}()

	wg.Wait()

	if len(errors) > 0 {
		return errors[0]
	}

	return nil
}

func (l Logic) CreateApiPage(ctx context.Context, table string) error {
	name := utils.Singular(table)
	res, err := l.viewer.Parse(ctx, apiPageTpl, g.Map{
		"resource": table,
		"name":     gstr.UcFirst(name),
		"version":  apiVersion,
	})
	if err != nil {
		return err
	}
	file := fmt.Sprintf("web/src/service/api/%s.ts", name)
	return gfile.PutContents(file, res)
}

func (l Logic) CreateListPage(ctx context.Context, fields model.Fields, table, path string) error {
	res, err := l.viewer.Parse(ctx, listPageTpl, g.Map{
		"table":  table,
		"fields": fields,
		"name":   gstr.UcFirst(utils.Singular(table)),
	})
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/index.vue", path)
	return gfile.PutContents(file, res)
}

func (l Logic) CreateFormPage(ctx context.Context, fields model.Fields, table, path string) error {
	res, err := l.viewer.Parse(ctx, formPageTpl, g.Map{
		"table":  table,
		"fields": fields,
		"name":   gstr.UcFirst(utils.Singular(table)),
	})
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/form.vue", path)
	return gfile.PutContents(file, res)
}

func (l Logic) CreateDataPage(ctx context.Context, fields model.Fields, table, path string) error {
	res, err := l.viewer.Parse(ctx, useDataTpl, g.Map{
		"table":  table,
		"fields": fields,
		"name":   gstr.UcFirst(utils.Singular(table)),
	})
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/useData.tsx", path)
	return gfile.PutContents(file, res)
}
