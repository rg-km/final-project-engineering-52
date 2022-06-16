package handlers

import (
	"net/http"
	"strconv"

	"go-scholarship/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type scholarHandler struct {
	scholarUseCase models.ScholarshipUseCase
}

// routes
func NewScholarshipHandler(r *gin.Engine, scholarUseCase models.ScholarshipUseCase) {
	handler := &scholarHandler{
		scholarUseCase: scholarUseCase,
	}
	// TODO: define routes
	r.GET("/api/scholarships", handler.fetch)
	r.GET("/api/scholarships/:id", handler.fetchById)
	r.POST("/api/scholarships", handler.create)
	r.PUT("/api/scholarships/:id", handler.update)
	r.DELETE("/api/scholarships/:id", handler.delete)



}

// fetch all scholarships
func (s *scholarHandler) fetch(c *gin.Context) {
	ctx := c.Request.Context()
	scholar, err := s.scholarUseCase.Fetch(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarships fetched",
		"data":    scholar,
	})
}

// fetch by id scholarship
func (s *scholarHandler) fetchById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	scholar, err := s.scholarUseCase.FetchById(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarship fetched",
		"data":    scholar,
	})
}

// create scholarship
func (s *scholarHandler) create(c *gin.Context) {
	ctx := c.Request.Context()
	scholar := models.ScholarRequest{}

	if err := c.ShouldBind(&scholar); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	res, err := s.scholarUseCase.Create(ctx, &scholar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarship created",
		"data":    res,
	})
}

// update scholarship
func (s *scholarHandler) update(c *gin.Context) {
	ctx := c.Request.Context()
	scholar := models.ScholarRequest{}
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	if err := c.ShouldBind(&scholar); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	res, err := s.scholarUseCase.Update(ctx, int64(idConv), &scholar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarship updated",
		"data":    res,
	})
}

// delete scholarship
func (s *scholarHandler) delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)

	err := s.scholarUseCase.Delete(ctx, int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarship deleted",
	})
}
