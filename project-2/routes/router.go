package routes

import (
	books "project-2/controllers/book"
	users "project-2/controllers/user"
	"project-2/middleware"
	m "project-2/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	// echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	// _ "github.com/swaggo/echo-swagger/example/docs"
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


    e.GET("/users", users.GetUserController)
	m.LogMiddleware(e)
	e.POST("/users",users.CreateUserController)
	e.POST("/login",users.LoginUserController)

    eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users",users.GetUserController)

    eBook := e.Group("/admin")
    eBook.Use(mid.BasicAuth(middleware.BasicAuthDB))
    eBook.POST("/book",books.CreateBook)
    eBook.GET("/book",books.GetBookController)
    eBook.GET("/book/:id",books.DetailBookController)
    eBook.PUT("/book/:id",books.UpdateBook)
    eBook.DELETE("/book/:id",books.DeleteBook)



	return e

}
