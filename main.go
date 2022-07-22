package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Numbers struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}
type Response struct {
	result interface{} `json:"result"`
}

func Addition(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}

	result := u.Num1 + u.Num2
	fmt.Println("The Addition of 2 numbers is=", result)

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

func SquareRoot(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := math.Sqrt(float64(u.Num1))
	fmt.Println("The Squareroot of  number is=", result)

	return c.JSON(http.StatusOK, result)
}
func Modulus(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}

	result := u.Num1 % u.Num2
	fmt.Println("The Modulus  of 2 numbers is=", result)

	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()

	e.POST("/addition", Addition)
	e.POST("/subtraction", Subtraction)
	e.POST("/multipication", Multiplication)
	e.POST("/division", Division)
	e.POST("/squareroot", SquareRoot)
	e.POST("/modulus", Modulus)

	e.Start(":4002")
}
