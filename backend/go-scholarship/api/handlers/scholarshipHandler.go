package handlers

import (
	"net/http"
	"strconv"

	"go-scholarship/api/handlers/middleware"
	"go-scholarship/api/models"

	"github.com/gin-gonic/gin"
)

type scholarHandler struct {
	scholarUseCase models.ScholarshipUseCase
}

// routes
func NewScholarshipHandler(r *gin.Engine, scholarUseCase models.ScholarshipUseCase) {
	handler := &scholarHandler{
		scholarUseCase: scholarUseCase,
	}

	// middleware
	m := middleware.InitMiddleware()
	auth := r.Group("/api").Use(m.JWTMiddleware())
	{
		auth.POST("/scholarships", handler.create)
		auth.PUT("/scholarships/:id", handler.update)
		auth.DELETE("/scholarships/:id", handler.delete)
	}

	// define routes
	r.GET("/api/scholarships", handler.fetch)
	r.GET("/api/scholarships/:id", handler.fetchById)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
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
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	scholar := models.ScholarRequest{}

	if err := c.ShouldBind(&scholar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
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
			"message": models.ItemNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scholarship deleted",
	})
}
