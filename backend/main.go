package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
)

func main() {
	app := pocketbase.New()

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "hello",
		Run: func(command *cobra.Command, args []string) {
			log.Println("Hello world!")
		},
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/golang/dojo",
			Handler: func(c echo.Context) error {
				return c.String(200, "Hello Gophers!")
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
