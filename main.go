package main

import (
	"echo-template/conf"
	"echo-template/db"
	"echo-template/routes"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
)


type CustomValidator struct{
	validator *validator.Validate
}

func (cv *CustomValidator)Validate(i interface{})error{
	return cv.validator.Struct(i)
}


//MAIN FUNCTION
func main() {

	//Init db connection
	db.Init()

	app := routes.Init()


	//CONFIGURATION FOR CORS
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins: "[]string{"*"}",
		// AllowedMethods: []string{"*"},
		// AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	//CORS MIDDLEWARE
	app.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	//INITIAL CONFIG ENV
	config := conf.GetConfig()

	//LOGGING MIDDLEWARE
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))

	//USING VALIDATOR
	app.Validator = &CustomValidator{validator: validator.New()}

	//VALIDATION ERROR HANDLING
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", 
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", 
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}
	
				break
			}
		}
	
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}


	

	app.Logger.Fatal(app.Start(config.PORT))
}