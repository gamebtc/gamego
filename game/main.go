package main

import (
	"flag"

	"local.com/abc/game/room"
)

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()

	router := NewGame()
	room.Start(name, router)
}
