package main

import (
	"github.com/jackkitte/youtube_manager_go/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8888"))
}
