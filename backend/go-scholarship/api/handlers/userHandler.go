package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-scholarship/api/handlers/middleware"
	"go-scholarship/api/models"
	"go-scholarship/utils/token"
)

type userHandler struct {
	userRepo models.UserRepository
}

// routes
func NewUserHandler(r *gin.Engine, userRepo models.UserRepository) {
	handler := &userHandler{
		userRepo: userRepo,
	}

	// auth middleware
	auth := r.Group("/api")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/users", handler.Fetch)
		auth.GET("/users/:id", handler.FetchById)
		auth.POST("/users", handler.Create)
		auth.PUT("/users/:id", handler.Update)
		auth.DELETE("/users/:id", handler.Delete) 
	}

	// should be public routes
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}

func errMessage(v validator.FieldError) string {
	m := fmt.Sprintf("error on field %s, condition: %s", v.Field(), v.ActualTag())

	return m
}

// login
func (u *userHandler) Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBind(&login); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	userLogin, err := u.userRepo.Login(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	// JWT
	token, _ := token.CreateToken(userLogin.Email)

	c.Request.Header.Set("Authorization", "Bearer "+token)

	// debug
	fmt.Println(c.Request.Header.Get("Authorization"))

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in",
		"token":   token,
		"data":    userLogin,
	})
}

// register
func (u *userHandler) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	userData, err := u.userRepo.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered",
		"data":    userData,
	})
}

// fetch users
func (u *userHandler) Fetch(c *gin.Context) {
	users, err := u.userRepo.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "users fetched",
		"data":    users,
	})
}

// fetch user by id
func (u *userHandler) FetchById(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}

	user, err := u.userRepo.FetchById(int64(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user fetched",
		"data":    user,
	})
}

// create user
func (u *userHandler) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	userData, err := u.userRepo.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"data":    userData,
	})
}

// update user
func (u *userHandler) Update(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}

	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			eM := errMessage(v)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": eM,
			})

			return
		}
	}

	userData, err := u.userRepo.Update(int64(idConv), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user":    userData,
	})
}

// delete user
func (u *userHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}

	if err := u.userRepo.Delete(int64(idConv)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
