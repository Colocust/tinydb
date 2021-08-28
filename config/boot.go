package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"tinydb/enum"
)

type Config struct {
	Addr         string `yaml:"addr"`
	MaxKeySize   uint32 `yaml:"max_key_size"`
	MaxValueSize uint32 `yaml:"max_value_size"`
	MaxClient    uint32 `yaml:"max_client"`
}

func NewConfig(fp string) (c *Config, res int) {
	data, err := ioutil.ReadFile(fp)

	if err != nil {
		log.Println("Error: Read config file error，The cause of the error is " + err.Error())
		res = enum.ERR
		return
	}

	c = new(Config)
	if err = yaml.Unmarshal(data, c); err != nil {
		log.Println("Error: Parsing config error，The cause of the error is " + err.Error())
		res = enum.ERR
		return
	}

	res = enum.OK
	return
}

func Load() (cfg *Config, ok int) {
	c := flag.String("c", "", "config file")
	flag.Parse()

	if *c == "" {
		log.Println("Warning: no config file specified, using the default config.")
		wd, _ := os.Getwd()
		*c = wd + "/config.yaml"
	}

	cfg, ok = NewConfig(*c)

	return
}
