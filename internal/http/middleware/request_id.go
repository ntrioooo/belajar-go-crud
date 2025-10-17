package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewString()
		c.Writer.Header().Set("X-Request-ID", id)
		c.Set("reqid", id)
		c.Next()
	}
}
