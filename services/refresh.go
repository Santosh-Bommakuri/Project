package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"Project/models"
)


func RefreshData(db *gorm.DB) {
	log.Println("Starting data refresh...")

	err := loadDataFromCSV(db)

	if err != nil {
		log.Println("Data refresh failed:", err)
	} else {
		log.Println("Data refresh completed successfully")
	}
}


func loadDataFromCSV(db *gorm.DB) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() 
	if err != nil {
		return fmt.Errorf("error reading header: %v", err)
	}

	
	tx := db.Begin()
	defer tx.Commit()

	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

	
		date, _ := time.Parse("2006-01-02", record[6])
		quantity, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)

		
		product := models.Product{
			ProductID: record[1],
			Name:      record[3],
			Category:  record[4],
			Region:    record[5],
		}
		tx.Where(models.Product{ProductID: product.ProductID}).
			Assign(product).
			FirstOrCreate(&product)

		
		customer := models.Customer{
			CustomerID: record[2],
			Name:       record[12],
			Email:      record[13],
			Address:    record[14],
		}
		tx.Where(models.Customer{CustomerID: customer.CustomerID}).
			Assign(customer).
			FirstOrCreate(&customer)

		
		order := models.Order{
			OrderID:       record[0],
			ProductID:     record[1],
			CustomerID:    record[2],
			DateOfSale:    date,
			QuantitySold:  quantity,
			UnitPrice:     unitPrice,
			Discount:      discount,
			ShippingCost:  shippingCost,
			PaymentMethod: record[11],
		}
		tx.Where(models.Order{OrderID: order.OrderID}).
			FirstOrCreate(&order)
	}

	return nil
}
