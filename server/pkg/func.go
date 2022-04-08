package pkg

import (
	"gf-admin/pkg/database"
	"github.com/gogf/gf/v2/container/gmap"
	"gorm.io/gorm"
)

var instances = gmap.NewStrAnyMap(true)

func DB() *gorm.DB {
	return database.GormDB
}
