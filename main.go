package main

import (
	. "github.com/torchcc/crank4go-router/crank_router"
)

func main() {
	app := CreateRouterApp(9000, 9070, 12439)
	app.Start()
	select {}
}
