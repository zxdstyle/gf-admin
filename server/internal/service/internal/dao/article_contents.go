// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/internal/service/internal/dao/internal"
)

// articleContentsDao is the data access object for table article_contents.
// You can define custom methods on it to extend its functionality as you wish.
type articleContentsDao struct {
	*internal.ArticleContentsDao
}

var (
	// ArticleContents is globally public accessible object for table article_contents operations.
	ArticleContents = articleContentsDao{
		internal.NewArticleContentsDao(),
	}
)

// Fill with you ideas below.
