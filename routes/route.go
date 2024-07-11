package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct{
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}


func Init() *echo.Echo {
	e:=echo.New()

	//ROUTE ENDPOINTS
	e.GET("/", func(ctx echo.Context)error{
		data := "This is index page"
		return ctx.String(http.StatusOK,data)
	})


	e.Any("/user", func(c echo.Context)(err error){
		u:=new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK,u)
	})

	e.POST("/post", func(c echo.Context)error{
		u :=new(User)
		if err:= c.Bind(u); err!=nil{
			return err
		}
		if err:= c.Validate(u); err!=nil{
			return err
		}
		return c.JSON(http.StatusOK, true)
	})


	return e
}