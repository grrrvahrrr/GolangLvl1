//Lesson 8 Homework

package config

import (
	"flag"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Name string
	Mode string
}

var (
	FlagName = flag.String("name", "", "names of people to greet")
	FlagMode = flag.String("mode", "", "hello or goodbye")
)

func (c *Configuration) ParseNames() []string {
	splitted := strings.Split(c.Name, ",")
	for i, s := range splitted {
		splitted[i] = strings.TrimSpace(s)
	}
	return splitted
}

func (c *Configuration) LoadConfig(file string) error {
	flag.Parse()
	err := godotenv.Load(file)
	if err != nil {
		return err
	}
	if *FlagName == "" {
		c.Name = os.Getenv("NAME")
	} else {
		c.Name = *FlagName
	}
	if *FlagMode == "" {
		c.Mode = os.Getenv("MODE")
	} else {
		c.Mode = *FlagMode
	}
	return nil
}

func Load(file string) (Configuration, error) {
	var c Configuration
	err := c.LoadConfig(file)
	if err != nil {
		return c, err
	}
	return c, nil
}
