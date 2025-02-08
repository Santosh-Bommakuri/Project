package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Project/models" // Import the models package
)

// DateRange request structure
type DateRange struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

// RevenueResponse structure
type RevenueResponse struct {
	TotalRevenue float64                 `json:"total_revenue"`
	ByProduct    []ProductRevenue        `json:"by_product,omitempty"`
	ByCategory   []CategoryRevenue       `json:"by_category,omitempty"`
	ByRegion     []RegionRevenue         `json:"by_region,omitempty"`
	Trends       []MonthlyRevenueTrend   `json:"trends,omitempty"`
}

type ProductRevenue struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Revenue   float64 `json:"revenue"`
}

type CategoryRevenue struct {
	Category string  `json:"category"`
	Revenue  float64 `json:"revenue"`
}

type RegionRevenue struct {
	Region  string  `json:"region"`
	Revenue float64 `json:"revenue"`
}

type MonthlyRevenueTrend struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}


func GetRevenueStats(c *gin.Context, db *gorm.DB) {
	var dateRange DateRange
	if err := c.ShouldBindJSON(&dateRange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, _ := time.Parse("2006-01-02", dateRange.StartDate)
	endDate, _ := time.Parse("2006-01-02", dateRange.EndDate)

	response := RevenueResponse{}

	// **Total Revenue**
	db.Model(&models.Order{}). // ✅ Fix: Use models.Order{}
		Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Select("SUM((unit_price * quantity_sold) * (1 - discount))").
		Scan(&response.TotalRevenue)

	// **Revenue by Product**
	db.Model(&models.Order{}). // ✅ Fix: Use models.Order{}
		Joins("JOIN products ON orders.product_id = products.product_id").
		Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Select("orders.product_id, products.name, SUM((unit_price * quantity_sold) * (1 - discount)) as revenue").
		Group("orders.product_id, products.name").
		Scan(&response.ByProduct)

	// **Revenue by Category**
	db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.product_id").
		Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Select("products.category, SUM((unit_price * quantity_sold) * (1 - discount)) as revenue").
		Group("products.category").
		Scan(&response.ByCategory)

	// **Revenue by Region**
	db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.product_id").
		Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Select("products.region, SUM((unit_price * quantity_sold) * (1 - discount)) as revenue").
		Group("products.region").
		Scan(&response.ByRegion)

	c.JSON(http.StatusOK, response)
}
