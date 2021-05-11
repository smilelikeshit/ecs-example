package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "os"

)

func main() {
  // Echo instance

  listen := ":"
  e := echo.New()


  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  

  // Start server
  e.Logger.Fatal(e.Start(listen + os.Getenv("PORT")))
}

// Handler
func hello(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : os.Getenv("HELLO"),
  })
  
}

