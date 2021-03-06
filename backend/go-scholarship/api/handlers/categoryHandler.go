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

	// define routes
	r.GET("/api/categories", handler.fetch)
	r.GET("/api/categories/:id", handler.fetchById)
	r.POST("/api/categories", handler.create)
	r.PUT("/api/categories/:id", handler.update)
	r.DELETE("/api/categories/:id", handler.delete)
}

// fetch all categories
func (ca *categoryHandler) fetch(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := ca.categoryRepo.Fetch(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, categories)
}

// fetch by id category
func (ca *categoryHandler) fetchById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	category, err := ca.categoryRepo.FetchById(ctx, int64(idConv))
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

// update category
func (ca *categoryHandler) update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	idConv, _ := strconv.Atoi(id)

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

	category, err := ca.categoryRepo.Update(ctx, int64(idConv), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, category)
}

// delete category
func (ca *categoryHandler) delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	idConv, _ := strconv.Atoi(id)

	err := ca.categoryRepo.Delete(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})
}
