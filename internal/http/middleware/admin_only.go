package middleware

import (
	"net/http"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// RequireAdmin: butuh user login & role=admin
func RequireAdmin(users ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidAny, ok := c.Get("userID")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User not logged in"})
			return
		}
		uid := uidAny.(uint)

		u, err := users.FindByID(c.Request.Context(), uid)
		if err != nil || u == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User not found"})
			return
		}
		if u.Role != domain.RoleAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "admin_only", "message": "Admin access required"})
			return
		}
		c.Next()
	}
}
