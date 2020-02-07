package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	"api-starter/db"
	"api-starter/modules/todos"
)

func main() {

	// init Gorm DB
	db.Init()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Home route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// add initModule function for all your modules
	todos.InitModule( e, db.DbManager() )

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}