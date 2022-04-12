package menu

import (
	"context"
	"gf-admin/app/logic/base"
	"gf-admin/app/model/system"
	"gf-admin/app/repository"
	"gf-admin/pkg/request"
)

type Logic struct {
	*base.Logic
}

func NewLogic() *Logic {
	return &Logic{
		Logic: base.NewLogic(repository.Menu()),
	}
}

func (l Logic) Tree(ctx context.Context, req request.ListReq, mos *system.Menus) error {
	if err := l.IndexAll(ctx, req, mos); err != nil {
		return err
	}

	refer := make(map[uint]*system.Menu)
	var tree system.Menus

	for i, mo := range *mos {
		refer[mo.GetID()] = (*mos)[i]
	}

	for i, mo := range *mos {
		pid := *mo.ParentId
		if pid == 0 {
			tree = append(tree, (*mos)[i])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = new(system.Menus)
				}
				*refer[pid].Children = append(*refer[pid].Children, (*mos)[i])
			}
		}
	}
	*mos = tree
	return nil
}
