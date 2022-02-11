package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page   int
	Size   int
	Offset int
	Total  int
}

func NewPagination(c *gin.Context) *Pagination {
	var p = &Pagination{}
	p.Page, p.Size = cast.ToInt(c.Query("pageNum")), cast.ToInt(c.Query("pageSize"))
	fmt.Print(p.Page)
	if p.Page == 0 || p.Size == 0 {
		p.Page, p.Size = 1, 10
	}
	p.Offset = (p.Page - 1) * p.Size
	return p
}

func (p *Pagination) GormPaginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.Size)
	}
}
