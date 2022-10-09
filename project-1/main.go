package main

import (
	"project-1/config"
	"project-1/routes"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func main(){
config.KonekDB()
e:=routes.New()
// e.Start(":3000")
   // Start server
e.Logger.Fatal(e.Start(":3000"))
}

