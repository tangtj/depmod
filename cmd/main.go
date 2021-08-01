package main

import (
	"flag"
	"github.com/tangtj/depmod/mod"
	"log"
	"os"
)

func main() {
	configPath := flag.String("c", "./config.json", "config file path")
	modPath := flag.String("m", "./go.mod", "go.mod file path")

	flag.Parse()

	ver, err := mod.Mod(*modPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cfg, err := mod.Config(*configPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	deponer := mod.NewDeponer(cfg)
	if err := deponer.Valid(ver); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
