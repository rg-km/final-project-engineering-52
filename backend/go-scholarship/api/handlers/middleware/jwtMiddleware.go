package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-scholarship/utils/token"
)

// header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := token.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}

// cookie
// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenStr, err := c.Cookie("token")
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		token.ValidateToken(tokenStr)

// 		c.Next()
// 	}
// }
