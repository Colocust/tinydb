package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Addr         string `yaml:"addr"`
	MaxKeySize   uint32 `yaml:"max_key_size"`
	MaxValueSize uint32 `yaml:"max_value_size"`
	MaxClient    uint32 `yaml:"max_client"`
}

func NewConfig(fp string) (c *Config) {
	data, err := ioutil.ReadFile(fp)

	if err != nil {
		log.Println("Error: Read config file error，The cause of the error is " + err.Error())
		os.Exit(1)
	}

	c = new(Config)
	if err = yaml.Unmarshal(data, c); err != nil {
		log.Println("Error: Parsing config error，The cause of the error is " + err.Error())
		os.Exit(1)
	}
	return
}

func Load() (cfg *Config) {
	c := flag.String("c", "", "config file")
	flag.Parse()

	if *c == "" {
		log.Println("Warning: no config file specified, using the default config.")
		wd, _ := os.Getwd()
		*c = wd + "/config.yaml"
	}

	cfg = NewConfig(*c)

	return
}
