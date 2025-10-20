package middleware

import (
	"strings"

	"belajar-go/pkg/jwtutil"

	"github.com/gin-gonic/gin"
)

// OptionalAuth: kalau ada token valid -> set userID ke context.
// Kalau tidak ada/invalid -> lanjut tanpa abort (viewer tetap 0).
func OptionalAuth(tm jwtutil.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		if v, err := c.Cookie("Authorization"); err == nil && v != "" {
			token = v
		}
		if token == "" {
			if auth := c.GetHeader("Authorization"); strings.HasPrefix(auth, "Bearer ") {
				token = strings.TrimPrefix(auth, "Bearer ")
			}
		}
		if token == "" {
			c.Next()
			return
		}

		if _, claims, err := tm.Parse(token); err == nil {
			if sub, ok := claims["sub"].(float64); ok && sub > 0 {
				c.Set("userID", uint(sub))
			}
		}
		c.Next()
	}
}
