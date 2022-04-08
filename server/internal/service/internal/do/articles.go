// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Articles is the golang structure of table articles for DAO operations like Where/Data.
type Articles struct {
	g.Meta      `orm:"table:articles, do:true"`
	Id          interface{} //
	Title       interface{} //
	CateId      interface{} //
	Summary     interface{} //
	OriginalUrl interface{} //
	IsPublish   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}