package repository

import (
	"Project/models"

	"gorm.io/gorm"
)


func UpsertProduct(db *gorm.DB, product models.Product) error {
	return db.Where(models.Product{ProductID: product.ProductID}).
		Assign(product).
		FirstOrCreate(&product).Error
}


func UpsertCustomer(db *gorm.DB, customer models.Customer) error {
	return db.Where(models.Customer{CustomerID: customer.CustomerID}).
		Assign(customer).
		FirstOrCreate(&customer).Error
}


func InsertOrder(db *gorm.DB, order models.Order) error {
	return db.Create(&order).Error
}
