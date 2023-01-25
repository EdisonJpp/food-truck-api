package main

import (
	"context"
	"fmt"
	CompanyModule "food-truck-api/api/company"
	ProductModule "food-truck-api/api/product"
	"food-truck-api/package/auth"
	"food-truck-api/package/company"
	"food-truck-api/package/product"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Databases struct {
	db        *mongo.Database
	productDB *mongo.Database
}

func main() {

	validate := validator.New()

	databases, cancel, err := databasesConnection()

	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	fmt.Println("Database connection success!")

	productCollection := databases.productDB.Collection("produtcs")
	productRepo := product.NewRepo(productCollection)
	productService := product.NewService(productRepo)

	companyCollection := databases.db.Collection("companies")
	companyRepo := company.NewRepo(companyCollection)
	companyService := company.NewService(companyRepo)

	authRepo := auth.NewRepo()
	authService := auth.NewService(authRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the food-truck-api shop!"))
	})

	api := app.Group("/api/v1")

	ProductModule.ProductRouter(api, productService)
	CompanyModule.CompanyRouter(api, companyService, authService, *validate)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databasesConnection() (*Databases, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017").SetServerSelectionTimeout(5*time.
		Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database("food_truck")
	productDB := client.Database("food_truck_products")

	databases := new(Databases)

	databases.db = db
	databases.productDB = productDB

	return databases, cancel, nil
}
