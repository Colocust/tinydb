package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tinydb/config"
	"tinydb/enum"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func main() {
	var (
		cfg *config.Config
		ok  int
	)

	if cfg, ok = loadConfig(); ok == enum.ERR {
		return
	}

	fmt.Println(cfg)
}

func loadConfig() (cfg *config.Config, ok int) {
	var c *string
	c = flag.String("c", "", "config file")
	flag.Parse()

	if *c == "" {
		log.Println("Warning: no config file specified, using the default config.")
		wd, _ := os.Getwd()
		*c = wd + "/config.yaml"
	}

	if cfg, ok = config.NewConfig(*c); ok == enum.OK {
		ok = cfg.Check()
	}

	return
}
