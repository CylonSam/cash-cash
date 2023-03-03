package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Outcome struct {
	ID          uint
	Description string
	Amount      float32
	Date        time.Time
}

type OutcomeResource struct{}

func (rs OutcomeResource) Routes(e *echo.Echo) {
	e.GET("/outcome", listOutcome)
	//e.GET("/outcome/:id", getOutcome)
	e.POST("/outcome", createOutcome)
	//e.PUT("/outcome/:id", updateOutcome)
	//e.DELETE("/outcome/:id", deleteOutcome)
}

func createOutcome(c echo.Context) error {
	outcome := new(Outcome)

	err := c.Bind(&outcome)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(outcome); err != nil {
		return err
	}

	incomeAlreadyExists := DB.Where(
		"description = ? AND EXTRACT(MONTH FROM date) = ?",
		outcome.Description,
		outcome.Date.Month()).First(&Income{})

	if incomeAlreadyExists.Error == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Outcome already exists")
	}

	createdIncome := DB.Create(&outcome)

	if createdIncome.Error != nil {
		return c.String(http.StatusInternalServerError, "oopsie!")
	}

	return c.String(http.StatusCreated, "Outcome created!")
}
