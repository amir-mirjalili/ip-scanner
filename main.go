package main

import (
	"github.com/amir-mirjalili/ip-scanner/cmd/api"
	"github.com/amir-mirjalili/ip-scanner/internal/db"
	"github.com/joho/godotenv"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	defer func() {
		if err := db.Close(database.DB); err != nil {
			log.Printf("⚠️ Failed to close database: %v", err)
		}
	}()

	if err := db.AutoMigrate(database); err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}

	server := api.NewServer(database)
	if err := server.App.Start(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
