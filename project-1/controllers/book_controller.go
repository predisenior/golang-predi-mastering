package controllers

import (
	"fmt"
	"net/http"
	"project-1/config"
	model "project-1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateBook(c echo.Context) error{
	book := model.Book{}

	c.Bind(&book)

	err:=config.DB.Create(&book).Error
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"book":book,
	})

}


func GetBookController(c echo.Context) error{
	var book []model.Book
	err:=config.DB.Find(&book).Error
	if err !=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"data": book,
	})
}

func DetailBookController(e echo.Context) error{
	var book []model.Book
	id,_:=strconv.Atoi(e.Param("id"))
	err:=config.DB.Find(&book,id).Error
	if err !=nil{
		return e.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return e.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"data": book,
	})

}

func UpdateBook(e echo.Context) error{
	// var book []model.Book
	book := model.Book{}
	e.Bind(&book)



	id,_:=strconv.Atoi(e.Param("id"))
	// err:=config.DB.Find(&book,id).Error
	book.Id = id
	// err:=config.DB.Save(&book,id).Error
	fmt.Println("id buku",id)
	err:=config.DB.Model(&model.Book{}).Where("id = ?", id).Updates(book).Error

	// err:=config.DB.Save(&book).Error

	if err !=nil{
		return e.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return e.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"data": book,
	})

}


func DeleteBook(c echo.Context) error{
	var book []model.Book
	id,_:=strconv.Atoi(c.Param("id"))
	err:=config.DB.Delete(&book,id).Error
	// err:=config.DB.Exec("DELETE FROM books WHERE id= ?",id).Error
	if err !=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"deleted":id,
		// "data": book,
	})



}
