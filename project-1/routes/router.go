package routes

import (
	"project-1/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2


func New() *echo.Echo{
	e :=echo.New()

//     Middleware
//    e.Use(middleware.Logger())
//    e.Use(middleware.Recover())
//    e.Use(middleware.CORS())

	e.POST("/book/",controllers.CreateBook)
	e.GET("/book/",controllers.GetBookController)
	e.GET("/book/:id",controllers.DetailBookController)
	e.PUT("/book/:id",controllers.UpdateBook)
	e.DELETE("/book/:id",controllers.DeleteBook)
    e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e

}
