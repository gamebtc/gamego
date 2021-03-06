package main

import (
	"flag"

	"local.com/abc/game/util"
)

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()
	defer util.PrintPanicStack()
	// open profiling
	config := InitConfig(name)
	Run(config)
}
