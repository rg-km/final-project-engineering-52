package middleware

// header
// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenStr := c.Request.Header.Get("Authorization")
// 		if tokenStr == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := token.ValidateToken(tokenStr)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user", claims)

// 		c.Next()
// 	}
// }
