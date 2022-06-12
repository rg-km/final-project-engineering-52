package handlers

import (
	"net/http"
	"strconv"

	"go-scholarship/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type categoryHandler struct {
	categoryRepo models.CategoryRepository
}

// routes
func NewCategoryHandler(r *gin.Engine, categoryRepo models.CategoryRepository) {
	handler := categoryHandler{categoryRepo}

	r.GET("/api/categories", handler.fetch)
	// TODO: define routes
	r.POST("/api/categories", handler.create)
}

// fetch all categories
func (repo *categoryHandler) fetch(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := repo.categoryRepo.Fetch(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, categories)
}

// fetch by id category
func (repo *categoryHandler) FetchById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	category, err := repo.categoryRepo.FetchById(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, category)
}

// create category
func (ca *categoryHandler) create(c *gin.Context) {
	ctx := c.Request.Context()
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	category, err := ca.categoryRepo.Create(ctx, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, category)
}

// TODO: Update

// TODO: Delete
