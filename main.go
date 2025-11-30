package main // This declares the main package, the entry point of the Go program

import (
	// Importing our initializer functions (LoadEnvVariables, ConnectToDb, SyncDatabase)
	"github.com/VINAYAK777CODER/go-auth/controllers"
	"github.com/VINAYAK777CODER/go-auth/initializers"

	// Importing the Gin framework for building HTTP APIs
	"github.com/gin-gonic/gin"
)

// init() runs automatically BEFORE main()
// Used to load environment variables, connect to the database,
// and sync the database schema.
func init() {
    initializers.LoadEnvVariables() // Load values from .env file
    initializers.ConnectToDb()      // Establish connection to PostgreSQL database
    initializers.SyncDatabase()     // Run GORM AutoMigrate to create/update tables
}

func main() {
    r := gin.Default() // Create a new Gin router with Logger + Recovery middleware

   r.POST("/signup",controllers.SignUp)

    r.Run() // Start the server on port 8080 by default
}
