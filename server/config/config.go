package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	ServerLogFilePath string `yaml:"log_file_path"`
	ListenAddress     string `yaml:"port"`
	JwtSecret         string `yaml:"jwt_secret"`

	DBDialect           string `yaml:"db_dialect"`
	DBUsername          string `yaml:"db_username"`
	DBPassword          string `yaml:"db_password"`
	DBName              string `yaml:"db_name"`
	DBHostname          string `yaml:"db_host"`
	DBParameters        string `yaml:"db_parameters"`
	DBPort              int    `yaml:"db_port"`
	DBProtocol          string `yaml:"db_protocol"`
	DBGormSingularTable bool   `yaml:"db_gorm_singular_table"`
	DBGormLogMode       bool   `yaml:"db_gorm_log_mode"`
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
		return fmt.Errorf("Could not read config: %s", err.Error())
	}

	if err := yaml.Unmarshal(content, conf); err != nil {
		return fmt.Errorf("Could not parse yaml config file '%s': %s", path, err.Error())
	}

	return nil
}

func NewConfig() *Config {
	return &Config{}
}
