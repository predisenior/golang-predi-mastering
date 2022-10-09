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
