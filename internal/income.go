package internal

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Income struct {
	gorm.Model
	Description string    `json:"description" validate:"required"`
	Amount      int       `json:"amount" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

type IncomeResource struct{}

func (rs IncomeResource) Routes(e *echo.Echo) {
	e.POST("/income", createIncome)
}

func createIncome(c echo.Context) error {
	income := new(Income)

	err := c.Bind(&income)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(income); err != nil {
		return err
	}

	incomeAlreadyExists := new(Income)
	DB.Where("description = ?", income.Description).First(&incomeAlreadyExists)

	if incomeAlreadyExists != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Income already exists")
	}

	createdIncome := DB.Create(&income)

	if createdIncome.Error != nil {
		return c.String(http.StatusInternalServerError, "oopsie!")
	}

	return c.String(http.StatusCreated, "user created!")
}
