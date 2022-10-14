package model


type Book struct{
	Id int			`json:"id" form:"id"`
	Title string	 	`json:"title" form:"title"`
	Author string	 	`json:"author" form:"author"`
	Publisher string	`json:"publisher" form:"publisher"`

}


type BooksResponse struct{
	ID int				`json:"id" form:"id"`
	Title string	 	`json:"title" form:"title"`
	Author string	 	`json:"author" form:"author"`
	Publisher string	`json:"publisher" form:"publisher"`
	Token string	 	`json:"token" form:"token"`

}



type User struct {
	ID int			`json:"id" form:"id"`
	Name string 	`json:"name" form:"name"`
	Email string 	`json:"email" form:"email"`
	Password string 	`json:"password" form:"password"`



}



type UserResponse struct {
	ID int			`json:"id" form:"id"`
	Name string 	`json:"name" form:"name"`
	Email string 	`json:"email" form:"email"`
	Token string 	`json:"token" form:"token"`

}
