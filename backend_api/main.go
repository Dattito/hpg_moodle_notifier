package main

import (
	"os"

	"github.com/Dattito/HMN_backend_api/app/endpoints"
	"github.com/Dattito/HMN_backend_api/app/utils"
	_ "github.com/Dattito/HMN_backend_api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func loadEnvs() {
	godotenv.Load(".env")
}

func setupRoutes(app *fiber.App) {
	app.Post("/v1/signalVerifications", endpoints.CreateSignalVerification)
	app.Post("/v1/assignments", endpoints.CreateAssignment)
	app.Delete("/v1/assignments", endpoints.DeleteAssignment)

	app.Post("/v1/moodleToken", endpoints.PostMoodleToken)

	app.Get("/health", endpoints.GetHeartbeat)
}

// @title HMN Backend API
// @version 1.0
// @description Register or Delete oneself from the HMN list
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 192.168.178.10:3000
// @BasePath /

func main() {
	loadEnvs()

	app := fiber.New()

	if os.Getenv("CORS") == "1" {
		app.Use(cors.New())
	}

	setupRoutes(app)
	utils.NewValidator()

	utils.StartServer(app)
}
