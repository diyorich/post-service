package handler

import (
	"github.com/diyorich/post-api/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"post-storage-service/internal/converter"
	"post-storage-service/internal/service"
	"strconv"
)

type handler struct {
	service service.PostService
}

func NewHandler(service service.PostService) *handler {
	return &handler{service: service}
}

func (h *handler) GetList(c *gin.Context) {
	p := pkg.GetPagination(c)
	data, err := h.service.GetList(c.Request.Context(), p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseErr(err))
		return
	}

	c.JSON(http.StatusOK, Response(converter.FromServiceToPostsJSON(data), p))
}

func (h *handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseErr(errors.New("invalid id passed")))
		return
	}

	post, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseErr(service.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, Response(converter.FromServiceToPostJSON(post), nil))
}
