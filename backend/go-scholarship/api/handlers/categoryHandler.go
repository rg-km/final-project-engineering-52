package handlers

import (
	"net/http"
	"strconv"

	"go-scholarship/api/models"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryRepo models.CategoryRepository
}

// routes
func NewCategoryHandler(r *gin.Engine, categoryRepo models.CategoryRepository) {
	handler := categoryHandler{categoryRepo}

	r.GET("/api/categories", handler.Fetch)
	// TODO: define routes
	
}

// fetch all categories
func (repo *categoryHandler) Fetch(c *gin.Context) {
	categories, err := repo.categoryRepo.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(200, categories)
}

// fetch by id category
func (repo *categoryHandler) FetchById(c *gin.Context) {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	category, err := repo.categoryRepo.FetchById(int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(200, category)
}

// TODO: Create

// TODO: Update

// TODO: Delete
