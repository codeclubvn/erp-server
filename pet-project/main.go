package main

import (
	"pet-project/config"
	"pet-project/route"
)

func main() {
	config.SetEnv()
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
