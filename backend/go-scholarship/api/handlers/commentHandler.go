package handlers

import (
	"net/http"
	"strconv"

	"go-scholarship/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type commentHandler struct {
	commentUseCase models.CommentUseCase
}

// routes
func NewCommentHandler(r *gin.Engine, commentRepo models.CommentRepository) {
	handler := commentHandler{commentRepo}

	// define routes
	r.GET("/api/comments", handler.fetch)
	r.GET("/api/comments/:id", handler.fetchById)
	r.POST("/api/comments", handler.create)
	r.PUT("/api/comments/:id", handler.update)
	r.DELETE("/api/comments/:id", handler.delete)
}

// fetch comments
func (co *commentHandler) fetch(c *gin.Context) {
	ctx := c.Request.Context()
	comments, err := co.commentUseCase.Fetch(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, comments)
}

// fetch by id comment
func (co *commentHandler) fetchById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	comment, err := co.commentUseCase.FetchById(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, comment)
}

// create comment
func (co *commentHandler) create(c *gin.Context) {
	ctx := c.Request.Context()
	var commentReq models.CommentRequest
	if err := c.ShouldBind(&commentReq); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	res, err := co.commentUseCase.Create(ctx, &commentReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, res)
}

// update comment
func (co *commentHandler) update(c *gin.Context) {
	ctx := c.Request.Context()
	var commentReq models.CommentRequest
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	if err := c.ShouldBind(&commentReq); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	res, err := co.commentUseCase.Update(ctx, int64(idConv), &commentReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, res)
}

// delete comment
func (co *commentHandler) delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	err := co.commentUseCase.Delete(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment deleted",
	})
}
