package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Application struct {
		Port string `yaml:"port"`
		Db   struct {
			Mysql struct {
				Host     string `yaml:"mysql"`
				Database string `yaml:"database"`
				Port     string `yaml:"port"`
				User     string `yaml:"user"`
				Password string `yaml:"password"`
			} `yaml:"mysql"`
		} `yaml:"db"`
	} `yaml:"application"`
}

var Config = &config{}

func Get() *config {

	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&Config); err != nil {
		panic("error[config]: can't decode config file")
	}
	return Config
}
