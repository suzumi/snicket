package handler

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := fasthttp.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if !c.Response().Committed() {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}

func Home(c echo.Context) error {
	return c.Render(fasthttp.StatusCreated, "home", "Hello")
}

func Snippet(c echo.Context) error {
	return c.Render(fasthttp.StatusCreated, "snippet", "Hello")
}

func SnippetCreate(c echo.Context) error {
	return c.Render(fasthttp.StatusCreated, "snippet_create", "Hello")
}
