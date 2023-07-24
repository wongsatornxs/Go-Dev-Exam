package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Cal struct {
	A int `json:"A"`
	B int `json:"B"`
}

type (
	Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
	SaveProductRespone struct{
		Message string `json:"message"`
		Data Product `json:"data"`
	}
)


var products []Product




func main() {
	
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.POST("/multiply", Calculate)
	e.GET("/products", getProducts)
	e.POST("/products", createProduct)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func Calculate(c echo.Context) error {
	cal := new(Cal)
	if err := c.Bind(cal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := cal.A * cal.B
	return c.JSON(http.StatusOK, map[string]int{
		"result": result,
	})
}
func getProducts(c echo.Context) error {

	return c.JSON(http.StatusOK, products)
}
func createProduct(c echo.Context) error {
	var newProduct Product
	err := c.Bind(&newProduct)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	responese := SaveProductRespone{
		Message: "add product successfully",
		Data: newProduct,
	}
	return c.JSON(http.StatusCreated, responese)
}
func init(){
	
	initProducts := []Product{{
		ID:    1,
		Name:  "สินค้าที่ 1",
		Price: 100,
	}, {

		ID:    2,
		Name:  "สินค้าที่ 2",
		Price: 200,
	}}
	products = append(products, initProducts...)
	
}
