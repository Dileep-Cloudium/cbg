package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cbglabs/nexusgate-pbmapi/internal/helper"
	"github.com/cbglabs/nexusgate-pbmapi/internal/router"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// @title nexusgate-pbmapi
// @version 1.0
// @description This is for the NexusGate PBM API.
// @description This solution is intended to act as a collection of CBG services that support the CBG PBM APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8889
// @BasePath /api/v1
// @schemes http
// swagger embed files
// gin-swagger middleware
// swagger embed files
func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		helper.ErrorPanic(err)
	}

	mainRouter := router.MountRoutes()

	server := &http.Server{
		Addr:           ":8889",
		Handler:        mainRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Sprintf("user#%s", uuid.New().String())
	server.ListenAndServe()
	fmt.Println("started server")

	// if serverErr != nil {
	// 	panic(err)
	// }
}
