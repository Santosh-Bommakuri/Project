package services

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"Project/models"
	"Project/repository"

	"gorm.io/gorm"
)


func LoadData(db *gorm.DB) error {

	file, err := os.Open("data.csv")
	if err != nil {
		return err
	}
	defer file.Close()


	reader := csv.NewReader(file)
	_, err = reader.Read() 
	if err != nil {
		return err
	}

	rowCount := 0
	batchSize := 1000
	products := []models.Product{}
	customers := []models.Customer{}
	orders := []models.Order{}

	
	for {
		record, err := reader.Read()
		if err != nil {
			break 
		}

		
		date, _ := time.Parse("2006-01-02", record[6])
		quantity, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)

		
		product := models.Product{ProductID: record[1], Name: record[3], Category: record[4], Region: record[5]}
		customer := models.Customer{CustomerID: record[2], Name: record[12], Email: record[13], Address: record[14]}
		order := models.Order{OrderID: record[0], ProductID: record[1], CustomerID: record[2], DateOfSale: date, QuantitySold: quantity, UnitPrice: unitPrice, Discount: discount, ShippingCost: shippingCost, PaymentMethod: record[11]}

		
		products = append(products, product)
		customers = append(customers, customer)
		orders = append(orders, order)

		rowCount++

		
		if len(products) >= batchSize {
			processBatch(db, products, customers, orders)
			products, customers, orders = nil, nil, nil
		}
	}

	if len(products) > 0 {
		processBatch(db, products, customers, orders)
	}

	log.Printf("Finished loading %d rows", rowCount)
	return nil
}

func processBatch(db *gorm.DB, products []models.Product, customers []models.Customer, orders []models.Order) {
	db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			repository.UpsertProduct(tx, product)
		}
		for _, customer := range customers {
			repository.UpsertCustomer(tx, customer)
		}
		for _, order := range orders {
			repository.InsertOrder(tx, order)
		}
		return nil
	})
}
