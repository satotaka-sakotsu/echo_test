package main

import (
  "example.com/echo_test/handler"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", handler.MainPage())

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}
