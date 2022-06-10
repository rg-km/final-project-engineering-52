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
func (repo *categoryHandler) FetchById (c *gin.Context)
	id := c.Param("id")
	idConv, err :strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message" :models.BadRequest,

		})
		return
	}
category, err := u.categoryRepo.FetchById(int64(idConv))
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": models.InternalServer,
	})
	return
}
// TODO: Create
func (u *categoryHandler) Create(c *gin.Context){
	var category models.User

	if err := c.ShouldBind(&category); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError,gin.H{
				"message": em,
			})
			return
		}
	}

	categoryData, err := u.categoryRepo.Create(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ 
			"message": models.InternalServer,
		})
		return
	}
	c. JSON(http.StatusOK, gin.H{
		"message": "category created",
		"data": categoryData,
	})
}


// TODO: Update
func (u *categoryHandler) Update(c *gin.Context) {
	id :=c.Param("id")
	idConv, err :=strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}
	var category models.Category
	
	if err := c.ShouldBind(&user); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)
			
			c.JSON(http.StatusInternalServerError, gin. H{
				"message": eM,
			})
			return
		}
	}

	categoryData, err := u.categoryRepo.Update(int64(idConv),&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "category update",
		"category": categoryData,
	})
}


// TODO: Delete
func (u *categoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": models.BadRequest,
		})
		return
	}
	if err := u.categoryRepo.Delete(int64(idConv)); err != nill {
		c.JSON(http.StatusInternalServerError,gin,H{
			"Message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "categories deleted",
	})
}
