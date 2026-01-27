package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenzchiro/desksetup-link-api/db"
	"github.com/kenzchiro/desksetup-link-api/handler"
	"github.com/kenzchiro/desksetup-link-api/repository"
	"github.com/kenzchiro/desksetup-link-api/service"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	database := db.Connect()
	defer database.Close()

	// Product setup
	productRepo := repository.NewProductRepository(database)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	// Highlight setup
	highlightRepo := repository.NewHighlightRepository(database)
	highlightSvc := service.NewHighlightService(highlightRepo)
	highlightHandler := handler.NewHighlightHandler(highlightSvc)

	// Router
	router := handler.NewRouter(productHandler, highlightHandler)

	port := os.Getenv("PORT")

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(router.Run(":" + port))
}
