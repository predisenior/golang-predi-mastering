package user

import (
	"net/http"
	"strconv"

	"project-2/config"
	model "project-2/models"

	"github.com/labstack/echo"
)



func GetUserController(c echo.Context) error {
	var users []model.User
	// DB.Find(&users)
    err:=config.DB.Find(&users).Error
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message":"success",
		"data":users,
})
}

func CreateUserController(c echo.Context) error{
	user:=model.User{}
	c.Bind(&user)

	err:=config.DB.Save(&user).Error
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : err.Error(),
		})
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"messages":"success create user",
		"user":user,
	})
}

func UserController(e echo.Context) error{
// fmt.Println("halo predator")
// return e.String(http.StatusOK,"hello world")
id, _:= strconv.Atoi(e.Param("id"))
user:=model.User{ID:id,Name:"predi",Email:"apasaja"}
search:=e.QueryParam("search")

return e.JSON(http.StatusOK,map[string]interface{}{
	"user":user,
	"search":search,
})
}


func LoginUserController(c echo.Context)error{
	user:=model.User{}
	c.Bind(&user)

	err:=config.DB.Where("email = ? AND password = ?", user.Email,user.Password).First(&user).Error
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : "gagal login",
			"error"	 : err.Error(),
		})
	}
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message" : "gagal login",
		})
	}


	return c.JSON(http.StatusOK,map[string]interface{}{
		"messages":"success create user",
	})
}
