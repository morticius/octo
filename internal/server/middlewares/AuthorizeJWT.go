package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/morticius/octo/internal/server/services"
	"log"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		jwtService := services.NewJWTAuthService()
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid token",
			})
			return
		}

		if token == nil || !token.Valid {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
