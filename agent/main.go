package main

import (
	"flag"

	"agent/conf"
	"agent/internal"

	"local.com/abc/game/util"
)

func main() {
	var name string
	flag.StringVar(&name, "conf", "app.yaml", "config file name")
	flag.Parse()
	defer util.PrintPanicStack()
	// open profiling
	config := conf.InitConfig(name)
	internal.Run(config)
}
