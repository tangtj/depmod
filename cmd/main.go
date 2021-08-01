package main

import (
	"depmod/mod"
	"flag"
)

func main() {
	configPath := flag.String("c", "./config.json", "config file path")
	modPath := flag.String("m", "./go.mod", "go.mod file path")

	flag.Parse()

	ver, err := mod.Mod(*modPath)
	if err != nil {
		panic(err)
	}

	cfg, err := mod.Config(*configPath)
	if err != nil {
		panic(err)
	}

	deponer := mod.NewDeponer(cfg)
	deponer.Valid(ver)
}
