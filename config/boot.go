package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"tinydb/enum"
)

type Config struct {
	Addr         string `yaml:"addr"`
	MaxKeySize   uint32 `yaml:"max_key_size"`
	MaxValueSize uint32 `yaml:"max_value_size"`
}

func NewConfig(filePath string) (c *Config, res int) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println("Error: Read config file error，The cause of the error is " + err.Error())
		res = enum.ERR
		return
	}

	c = new(Config)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Println("Error: Parsing config error，The cause of the error is " + err.Error())
		res = enum.ERR
		return
	}

	res = enum.OK
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
