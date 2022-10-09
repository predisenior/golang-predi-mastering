package routes

import (
	"project-1/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo{
	e :=echo.New()

    // Middleware
   e.Use(middleware.Logger())
   e.Use(middleware.Recover())
   e.Use(middleware.CORS())

	e.POST("/book/",controllers.CreateBook)
	e.GET("/book/",controllers.GetBookController)
	e.GET("/book/:id",controllers.DetailBookController)
	e.PUT("/book/:id",controllers.UpdateBook)
	e.DELETE("/book/:id",controllers.DeleteBook)


	return e

}
