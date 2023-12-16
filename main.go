package main

import (
	"html/template"
	"portfolio/handler"
	"portfolio/render"

	"github.com/labstack/echo/v4"
)

func main() {
	r := &render.Renderer{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = r
	e.Static("/dist", "dist")

	e.GET("/", handler.Home)

	e.Logger.Fatal(e.Start(":8080"))
}
