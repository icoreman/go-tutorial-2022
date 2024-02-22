package main

import (
	"fmt"
	rookout "github.com/Rookout/GoSDK"
	"github.com/Rookout/go-tutorial-2022/pkg/todos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	rookout.Start(rookout.RookOptions{
		Token: "36f764f494198b33b8a8476c26aeb18a70007a5fe5918fe2f82769d84b0d6b77",
		Labels: map[string]string{"env": "dev"},

	})
	port := "8080"

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	todos.RegisterEndPoints(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
