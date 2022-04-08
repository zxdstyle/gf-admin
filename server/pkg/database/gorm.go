package database

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB = ConnectGorm()

func ConnectGorm() *gorm.DB {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		g.Cfg().MustGet(ctx, "database.default.user", "root").String(),
		g.Cfg().MustGet(ctx, "database.default.pass").String(),
		g.Cfg().MustGet(ctx, "database.default.host", "127.0.0.1").String(),
		g.Cfg().MustGet(ctx, "database.default.port", 3306).Int(),
		g.Cfg().MustGet(ctx, "database.default.name", "gf-admin").String(),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if g.Cfg().MustGet(ctx, "database.default.debug").Bool() {
		db = db.Debug()
	}

	return db
}
