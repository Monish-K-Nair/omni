package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
)

// Generate a unique ID and attach it to request

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
