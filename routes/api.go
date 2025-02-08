
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Project/services"
)

// SetupRoutes initializes all API endpoints
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	
	r.POST("/api/revenue", func(c *gin.Context) { services.GetRevenueStats(c, db) })
	
}
