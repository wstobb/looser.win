package main

import (
	"embed"

	"github.com/wstobb/looser.win/internal/app"
)

//go:embed static
var static embed.FS

//go:embed templates
var templates embed.FS

func main() {
	if err := app.New(static, templates).Start(); err != nil {
		panic(err)
	}
}
