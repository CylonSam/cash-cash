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
	e.GET("/income", listIncome)
  e.GET("/income/:id", getIncome)
	e.POST("/income", createIncome)
}

func getIncome(c echo.Context) error {
  ID := c.Param("id")
  income := new(Income)
  DB.First(&income, ID)

  return c.JSON(http.StatusOK, income)
}

func listIncome(c echo.Context) error {
	incomes := new([]Income)
	DB.Table("incomes").Find(&incomes)
  
	return c.JSON(http.StatusOK, incomes)
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

	incomeAlreadyExists := DB.Where("description = ?", income.Description).First(&Income{})

	if incomeAlreadyExists.Error == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Income already exists")
	}

	createdIncome := DB.Create(&income)

	if createdIncome.Error != nil {
		return c.String(http.StatusInternalServerError, "oopsie!")
	}

	return c.String(http.StatusCreated, "user created!")
}
