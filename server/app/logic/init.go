package logic

import (
	"gf-admin/app/logic/system/menu"
	"gf-admin/app/logic/system/permission"
	"gf-admin/app/logic/system/role"
	"gf-admin/app/logic/system/scaffold"
	"gf-admin/app/logic/user"
	"github.com/gogf/gf/v2/container/gmap"
)

var logic = gmap.NewStrAnyMap(true)

func Scaffold() *scaffold.Logic {
	return logic.GetOrSetFuncLock("logic.scaffold", func() interface{} {
		return scaffold.NewLogic()
	}).(*scaffold.Logic)
}

func User() *user.Logic {
	return logic.GetOrSetFuncLock("logic.user", func() interface{} {
		return user.NewLogic()
	}).(*user.Logic)
}

func Role() *role.Logic {
	return logic.GetOrSetFuncLock("logic.role", func() interface{} {
		return role.NewLogic()
	}).(*role.Logic)
}

func Permission() *permission.Logic {
	return logic.GetOrSetFuncLock("logic.permission", func() interface{} {
		return permission.NewLogic()
	}).(*permission.Logic)
}

func Menu() *menu.Logic {
	return logic.GetOrSetFuncLock("logic.menu", func() interface{} {
		return menu.NewLogic()
	}).(*menu.Logic)
}
