// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Chapters is the golang structure for table chapters.
type Chapters struct {
	Id        uint        `json:"id"        ` //
	BookId    int         `json:"bookId"    ` // 所属文档
	Name      string      `json:"name"      ` // 章节名称
	SortNum   uint        `json:"sortNum"   ` // 排序值，值越大越靠前
	IsActive  uint        `json:"isActive"  ` // 是否启用
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}