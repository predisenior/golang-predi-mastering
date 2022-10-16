package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "project-3/service"


type Resolver struct{
        userService service.UserService
        bookService service.BookService
}
