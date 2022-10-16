package middleware

import (
	"project-3/config"
	// model "project-3/models"
	model "project-3/graph/model"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool,error) {
	var user model.User
	err :=config.DB.Where("email = ? AND password = ?",username, password).First(&user).Error
	if (err!=nil){
		return false, err
	}
	return true, nil
	}
