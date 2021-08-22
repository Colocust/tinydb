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

	if cfg, ok = NewConfig(*c); ok == enum.OK {
		ok = cfg.Check()
	}

	return
}

func (cfg *Config) Check() (res int) {
	if cfg.Addr == "" {
		res = enum.ERR
		log.Println("Error: The addr config error")
		return
	}

	if cfg.MaxKeySize == 0 {
		res = enum.ERR
		log.Println("Error: The max_key_size config error")
		return
	}

	if cfg.MaxValueSize == 0 {
		res = enum.ERR
		log.Println("Error: The max_value_size config error")
		return
	}

	res = enum.OK
	return
}
