package main

import (
	"os"

	"github.com/singgihdwindaru/goSimpleBlog.Api/core/app"
	_ "github.com/singgihdwindaru/goSimpleBlog.Api/core/app"
)

func main() {
	app := app.App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8010")
}
