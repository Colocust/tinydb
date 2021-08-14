package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tinydb/config"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		return
	}
	fmt.Println(cfg)
}

func loadConfig() (cfg *config.Config, err error) {
	var c *string
	c = flag.String("c", "", "config file")
	flag.Parse()

	if *c == "" {
		log.Println("Warning: no config file specified, using the default config.")
		wd, _ := os.Getwd()
		*c = wd + "/config.yaml"
	}
	cfg, err = config.NewConfig(*c)
	return
}
