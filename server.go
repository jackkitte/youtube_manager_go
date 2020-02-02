package main

import (
	"github.com/jackkitte/youtube_manager_go/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8888"))
}
