package main

import (
	"cash-cash/internal"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	internal.ConnectToDatabase()

	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}
	internal.IncomeResource{}.Routes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
