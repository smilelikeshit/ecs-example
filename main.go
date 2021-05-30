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
  e.GET("/v1/video", v1video)
  e.GET("/v1/phonebook", v1phonebook)
  e.GET("/v2/video", v2video)
  e.GET("/v2/phonebook", v2phonebook)
  

  // Start server
  e.Logger.Fatal(e.Start(listen + os.Getenv("PORT")))
}

// Handler
func hello(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : os.Getenv("HELLO"),
  })
  
}

func v1video(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : "this is service v1/video",
  })
  
}

func v1phonebook(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : "this is service v1/phonebook",
  })
  
}


func v2video(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : "this is service v2/video",
  })
  
}

func v2phonebook(c echo.Context) error {
  
  return c.JSON(http.StatusOK, map[string]string{
    "message" : "this is service v2/phonebook",
  })
  
}

