package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-scholarship/api/models"
)

type categoryHandler struct {
	categoryRepo models.CategoryRepository
}

func NewCategoryHandler(r *gin.Engine, categoryRepo models.CategoryRepository) {
	handler := categoryHandler{categoryRepo}

	r.GET("/api/categories", handler.fetch)
}

// Fetch
func (repo *categoryHandler) fetch(c *gin.Context) {
	categories, err := repo.categoryRepo.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(200, categories)
}

// TODO: FetchById

// TODO: Create

// TODO: Update

// TODO: Delete
