package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-scholarship/api/handlers/middleware"
	"go-scholarship/api/models"
	"go-scholarship/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		auth.GET("/users", handler.fetch)
		auth.GET("/users/:id", handler.fetchById)
		auth.POST("/users", handler.create)
		auth.PUT("/users/:id", handler.update)
		auth.DELETE("/users/:id", handler.delete)
	}

	// should be public routes
	r.POST("/login", handler.login)
	r.POST("/register", handler.register)
}

func errMessage(v validator.FieldError) string {
	m := fmt.Sprintf("error on field %s, condition: %s", v.Field(), v.ActualTag())

	return m
}

// login
func (u *userHandler) login(c *gin.Context) {
	ctx := c.Request.Context()
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

	userLogin, err := u.userRepo.Login(ctx, &login)
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
func (u *userHandler) register(c *gin.Context) {
	ctx := c.Request.Context()
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

	userData, err := u.userRepo.Register(ctx, &user)
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
func (u *userHandler) fetch(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := u.userRepo.Fetch(ctx)
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
func (u *userHandler) fetchById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}

	user, err := u.userRepo.FetchById(ctx, int64(idConv))
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
func (u *userHandler) create(c *gin.Context) {
	ctx := c.Request.Context()
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

	userData, err := u.userRepo.Create(ctx, &user)
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
func (u *userHandler) update(c *gin.Context) {
	ctx := c.Request.Context()
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

	userData, err := u.userRepo.Update(ctx, int64(idConv), &user)
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
func (u *userHandler) delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": models.BadRequest,
		})
		return
	}

	if err := u.userRepo.Delete(ctx, int64(idConv)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": models.InternalServer,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
