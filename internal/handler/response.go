package handler

import (
	"github.com/diyorich/post-api/pkg"
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func Response(data interface{}, pagination *pkg.Pagination) gin.H {
	var meta Meta
	if pagination != nil {
		meta = Meta{
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
			Total:  pagination.Total,
		}
	}

	return gin.H{
		"data": data,
		"meta": meta,
	}
}

func ResponseErr(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
