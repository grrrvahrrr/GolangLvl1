package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	NAME string `yaml:"name" json:"name"`
	MODE string `yaml:"mode" json:"mode"`
}

func (c *Configuration) ParseNames() []string {
	splitted := strings.Split(c.NAME, ",")
	for i, s := range splitted {
		splitted[i] = strings.TrimSpace(s)
	}
	return splitted
}

func (c *Configuration) LoadConfig(name string, mode string) error {
	err := godotenv.Load("config_example.env")
	if err != nil {
		return err
	}
	if name == "" {
		c.NAME = os.Getenv("NAME")
	} else {
		c.NAME = name
	}
	if mode == "" {
		c.MODE = os.Getenv("MODE")
	} else {
		c.MODE = mode
	}
	return nil
}

func (c *Configuration) ConfigFromJsonYaml(file string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if strings.Contains(file, ".json") {
		if err = json.Unmarshal(contents, &c); err != nil {
			return err
		}
	}
	if strings.Contains(file, ".yaml") {
		if err = yaml.Unmarshal(contents, &c); err != nil {
			return err
		}
	}
	return nil
}
