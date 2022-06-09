package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-scholarship/api/models"
)

type scholarHandler struct {
	scholarRepo models.ScholarshipRepository
}

// routes
func NewScholarshipHandler(r *gin.Engine, scholarRepo models.ScholarshipRepository) {
	handler := &scholarHandler{
		scholarRepo: scholarRepo,
	}

	r.GET("/api/scholarships", handler.Fetch)
}

// Fetch
func (s *scholarHandler) Fetch(c *gin.Context) {
	scholar, err := s.scholarRepo.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    scholar,
	})
}
