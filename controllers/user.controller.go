package controllers

import (
	"echo-template/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllUser(c echo.Context)error{
	result,err := models.FetchAllUser()
	if err !=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostUser(c echo.Context)error{
	name := c.FormValue("name")
	city := c.FormValue("city")

	result, err:= models.PostUser(name, city)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}


	return c.JSON(http.StatusOK, result)
}


func UpdateUser(c echo.Context)error{
	id :=c.Param("id")
	name :=c.FormValue("name")
	city :=c.FormValue("city")

	convert_id,err := strconv.Atoi(id)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err:=models.UpdateUser(convert_id,name,city)
	if err !=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context)error{
	id:=c.Param("id")

	convert_id,err:=strconv.Atoi(id)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	result, err:= models.DeleteUser(convert_id)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:":err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}