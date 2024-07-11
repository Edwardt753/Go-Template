package controllers

import (
	"echo-template/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllUser(c echo.Context)error{
	result,err := models.FetchAllUser()
	if err !=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}