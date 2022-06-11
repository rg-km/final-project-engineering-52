package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-scholarship/api/models"
)

type commentHandler struct {
	commentRepo models.CommentRepository
}

// routes
func NewCommentHandler(r *gin.Engine, commentRepo models.CommentRepository) {
	handler := commentHandler{commentRepo}

	r.GET("/api/comments", handler.fetch)
}

// Fetch
func (repo *commentHandler) fetch(c *gin.Context) {
	comments, err := repo.commentRepo.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(200, comments)
}

// TODO: FetchById

// TODO: Create

// TODO: Update

// TODO: Delete
