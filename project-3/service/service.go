package service

import (
	"errors"
	"project-3/config"

	// model "project-3/models"
	model "project-3/graph/model"
)

type UserService struct{}
type BookService struct{}

// users

func (u *UserService) CreateUser(input model.UserInput) model.User{
    var newUser model.User = model.User{
        // ID: fmt.Sprintf("%d",rand.Int()),
        Name: input.Name,
        Email: input.Email,
        Password: input.Password,
    }
    config.DB.Create(&newUser)

    return newUser
}
func (u *UserService) GetAllUsers() []*model.User{
    var users []*model.User = []*model.User{}
    config.DB.Find(&users)
    return  users
}


func (u *UserService) GetUserByID(id string) (model.User, error){
    var user model.User
    result:=config.DB.First(&user,"id = ?",id)
    if result.RowsAffected == 0 {
        return model.User{},errors.New(" data tidak ditemukan ")
    }
    return user, nil

}
func (u *UserService) UpdateUser(input *model.UserInput, id int) error{
    user := model.User{
        // ID:        id,
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
    }
     err:=config.DB.Model(&user).Where("id = ?", id).Updates(user).Error
    return err
}



func (u *UserService) DeleteUser(id int) error{
    var user model.User
    err := config.DB.Delete(user, id).Error
    return err
}

//end users


func (u *BookService) GetAllBooks() []*model.Book{
    var books []*model.Book = []*model.Book{}
    config.DB.Find(&books)
    return  books
}

func (u *BookService) GetBookByID(id string) (model.Book, error){
    var book model.Book
    result:=config.DB.First(&book,"id = ?",id)
    if result.RowsAffected == 0 {
        return model.Book{},errors.New(" data buku tidak ditemukan ")
    }
    return book, nil
}

func (u *BookService) CreateBook(input model.BookInput) model.Book{
    var newBook model.Book = model.Book{
        // ID: fmt.Sprintf("%d",rand.Int()),
        Title: input.Title,
        Author: input.Author,
        Publisher: input.Publisher,
    }
    config.DB.Create(&newBook)
    return newBook
}

func (u *UserService) UpdateBook(input *model.BookInput, id int) error{
    book := model.Book{
        // ID:        id,
        Title:     input.Title,
        Author:    input.Author,
        Publisher: input.Publisher,
    }
     err:=config.DB.Model(&book).Where("id = ?", id).Updates(book).Error
    return err
}



func (u *UserService) DeleteBook(id int) error{
    var book model.Book
    err := config.DB.Delete(book, id).Error
    return err
}
