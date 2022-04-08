package repository

import (
	"gf-admin/app/repository/permission"
	"gf-admin/app/repository/role"
	"gf-admin/app/repository/scaffold"
	"gf-admin/app/repository/user"
	"github.com/gogf/gf/v2/container/gmap"
)

var repositories = gmap.NewStrAnyMap(true)

func Scaffold() scaffold.Repository {
	return repositories.GetOrSetFuncLock("repository.scaffold", func() interface{} {
		return scaffold.NewDbRepository()
	}).(scaffold.Repository)
}

func User() user.Repository {
	return repositories.GetOrSetFuncLock("repository.users", func() interface{} {
		return user.NewDbRepository()
	}).(user.Repository)
}

func Role() role.Repository {
	return repositories.GetOrSetFuncLock("repository.roles", func() interface{} {
		return role.NewDbRepository()
	}).(role.Repository)
}

func Permission() permission.Repository {
	return repositories.GetOrSetFuncLock("repository.permissions", func() interface{} {
		return permission.NewDbRepository()
	}).(permission.Repository)
}
