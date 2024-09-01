package main

import (
	"log"
	"os"
	"traffy-mock-crud/configuration"
	ds "traffy-mock-crud/domain/datasources"
	repo "traffy-mock-crud/domain/repositories"
	gw "traffy-mock-crud/src/gateways"
	"traffy-mock-crud/src/middlewares"
	sv "traffy-mock-crud/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	postgresql := ds.NewPostgreSQL(10)
	awsS3, err := sv.NewS3Uploader(os.Getenv("AWS_S3_BUCKET_NAME"), os.Getenv("AWS_REGION"))

	if err != nil {
		log.Fatal(err)
	}

	reportRepo := repo.NewReportsRepository(postgresql)

	sv1 := sv.NewReportsService(reportRepo, awsS3)

	gw.NewHTTPGateway(app, sv1)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
