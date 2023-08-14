package main

import (
	"erp-server/conf"
	"erp-server/route"
)

func main() {
	conf.SetEnv()
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
