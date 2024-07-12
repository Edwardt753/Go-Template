package routes

import (
	"echo-template/controllers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct{
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}




func Init() *echo.Echo {
	e:=echo.New()

	// //USING VALIDATOR
	e.Validator = &CustomValidator{validator: validator.New()}


	//ROUTE ENDPOINTS
	e.GET("/", func(ctx echo.Context)error{
		data := "This is index page"
		return ctx.String(http.StatusOK,data)
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


	e.GET("/user", controllers.FetchAllUser)

	e.POST("/user", controllers.PostUser)

	e.PUT("/user/:id", controllers.UpdateUser)

	e.DELETE("user/:id", controllers.DeleteUser)




	return e
}