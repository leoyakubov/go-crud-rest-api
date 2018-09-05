package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	ServerLogFilePath string `yaml:"log_file_path"`
	ListenAddress     string `yaml:"port"`
}

func (conf *Config) Read(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("No config file found by the path: " + path)
	}

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read config: %s", err.Error())
	}

	if err := yaml.Unmarshal(content, conf); err != nil {
		return fmt.Errorf("could not parse yaml config file '%s': %s", path, err.Error())
	}

	log.Println("Server config: ", *conf)

	return nil
}

func NewConfig() *Config {
	return &Config{}
}
