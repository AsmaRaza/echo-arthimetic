package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Numbers struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}
type Response struct {
	result int `json:"result"`
}

func Addition(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 + u.Num2
	fmt.Println("The Addirtion of 2 numbers is=", result)

	return c.JSON(http.StatusOK, result)
}
func Subtraction(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 - u.Num2
	fmt.Println("The Subtraction of 2 numbers is=", result)

	return c.JSON(http.StatusOK, result)
}
func Multiplication(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 * u.Num2
	fmt.Println("The Multiplication of 2 numbers is=", result)

	return c.JSON(http.StatusOK, result)
}
func Division(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 / u.Num2
	fmt.Println("The Division of 2 numbers is=", result)

	return c.JSON(http.StatusOK, result)
}
func main() {
	e := echo.New()

	e.POST("/", Addition)
	e.POST("/", Subtraction)
	e.POST("/", Multiplication)
	e.POST("/", Division)

	e.Start(":3008")
}
