package main

import (
	"log"
	"time"

	"Project/config"
	"Project/routes"
	"Project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}


	if err := services.LoadData(db); err != nil {
		log.Fatal("Failed to load data:", err)
	}

	services.RefreshData(db)

	// Start a background Goroutine for periodic refresh every 24 hours
	go func() {
		for {
			now := time.Now()
			nextRun := now.Truncate(24 * time.Hour).Add(24 * time.Hour) 
			time.Sleep(time.Until(nextRun))                             

			log.Println("Running scheduled data refresh...")
			services.RefreshData(db)
		}
	}()

	
	r := gin.Default()

	
	routes.SetupRoutes(r, db)

	
	r.Run(":8081")

}
