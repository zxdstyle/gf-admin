// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleContents is the golang structure of table article_contents for DAO operations like Where/Data.
type ArticleContents struct {
	g.Meta       `orm:"table:article_contents, do:true"`
	Id           interface{} //
	ArticleId    interface{} //
	MarkdownText interface{} //
	HtmlText     interface{} //
}
