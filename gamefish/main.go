package main

import (
	"flag"

	"gamezjh/internal"
	"local.com/abc/game/room"
)

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()

	router := internal.NewGame()
	room.Start(name, router)
}
