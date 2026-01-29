package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenzchiro/desksetup-link-api/db"
	"github.com/kenzchiro/desksetup-link-api/handler"
	categoryRepository "github.com/kenzchiro/desksetup-link-api/repositories/category"
	highlightRepository "github.com/kenzchiro/desksetup-link-api/repositories/highlight"
	productRepository "github.com/kenzchiro/desksetup-link-api/repositories/product"
	"github.com/kenzchiro/desksetup-link-api/services/category"
	"github.com/kenzchiro/desksetup-link-api/services/highlight"
	"github.com/kenzchiro/desksetup-link-api/services/product"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	database := db.Connect()
	defer database.Close()

	// Product setup
	productRepo := productRepository.NewProductRepository(database)
	productSvc := product.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	// Highlight setup
	highlightRepo := highlightRepository.NewHighlightRepository(database)
	highlightSvc := highlight.NewHighlightService(highlightRepo, productRepo)
	highlightHandler := handler.NewHighlightHandler(highlightSvc)

	// Category setup
	categoryRepo := categoryRepository.NewCategoryRepository(database)
	categorySvc := category.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categorySvc)

	// Router
	router := handler.NewRouter(productHandler, highlightHandler, categoryHandler)

	port := os.Getenv("PORT")

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(router.Run(":" + port))
}
