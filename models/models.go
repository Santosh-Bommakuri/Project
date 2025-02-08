
package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID string `gorm:"uniqueIndex"`
	Name      string
	Category  string
	Region    string
	Orders    []Order `gorm:"foreignKey:ProductID;references:ProductID"`
}

type Customer struct {
	ID         uint   `gorm:"primaryKey"`
	CustomerID string `gorm:"uniqueIndex"`
	Name       string
	Email      string
	Address    string
	Orders     []Order `gorm:"foreignKey:CustomerID;references:CustomerID"`
}

type Order struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       string    `gorm:"uniqueIndex"`
	ProductID     string
	CustomerID    string
	DateOfSale    time.Time
	QuantitySold  int
	UnitPrice     float64
	Discount      float64
	ShippingCost  float64
	PaymentMethod string
}
