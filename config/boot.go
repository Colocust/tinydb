package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Addr         string `yaml:"addr"`
	MaxKeySize   uint32 `yaml:"max_key_size"`
	MaxValueSize uint32 `yaml:"max_value_size"`
}

func NewConfig(filePath string) (c *Config, err error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println("Error: Read config file error，The cause of the error is " + err.Error())
		return
	}

	c = new(Config)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Println("Error: Parsing config error，The cause of the error is " + err.Error())
		return
	}
	return
}
