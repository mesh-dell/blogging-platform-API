package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/mesh-dell/blogging-platform-API/internal/blog"
	"github.com/mesh-dell/blogging-platform-API/internal/blog/service"
)

type BlogHandler struct {
	svc service.BlogPostService
}

func NewBlogHandler(svc service.BlogPostService) *BlogHandler {
	return &BlogHandler{
		svc: svc,
	}
}

func (h *BlogHandler) GetAll(c *gin.Context) {
	posts, err := h.svc.GetBlogs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func (h *BlogHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	post, err := h.svc.GetBlogById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, post)
}

func (h *BlogHandler) Create(c *gin.Context) {
	var req model.BlogPostRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svc.CreateBlog(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, res)
}

func (h *BlogHandler) Update(c *gin.Context) {
	var req model.BlogPostRequest
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := h.svc.UpdateBlog(id, req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, post)
}

func (h *BlogHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.DeleteBlog(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
