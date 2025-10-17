package middleware

import (
	"net/http"
	"strings"

	"belajar-go/pkg/jwtutil"

	"github.com/gin-gonic/gin"
)

// RequireAuth memverifikasi JWT, set userID ke context, lalu next.
// Pakai cookie "Authorization" atau header "Authorization: Bearer <token>".
func RequireAuth(tm jwtutil.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari cookie
		token := ""
		if v, err := c.Cookie("Authorization"); err == nil && v != "" {
			token = v
		}
		// Fallback: header Authorization
		if token == "" {
			auth := c.GetHeader("Authorization")
			if strings.HasPrefix(auth, "Bearer ") {
				token = strings.TrimPrefix(auth, "Bearer ")
			}
		}

		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, claims, err := tm.Parse(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(float64) // HS256 + MapClaims -> float64
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", uint(sub))
		c.Next()
	}
}
