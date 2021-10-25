//Lesson 8 Homework

package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Configuration struct {
	NAME string
	MODE string
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
