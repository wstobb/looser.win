package main

import (
	"github.com/wstobb/looser.win/internal/app"
)

func main() {
	if err := app.New().Start(); err != nil {
		panic(err)
	}
}
