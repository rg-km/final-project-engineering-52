package repository

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	bodyReader := strings.NewReader(`{
		"email": "test@gmail.com",
		"password": "test123",
	}`)

	req := httptest.NewRequest("POST", "http://localhost:8080/login", bodyReader)

	defer req.Body.Close()

	r.Run()

	assert.Equal(t, http.StatusOK, w.Code)
}
