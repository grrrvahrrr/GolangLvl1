package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func ParseNames(raw string) []string {
	splitted := strings.Split(raw, ",")
	for i, s := range splitted {
		splitted[i] = strings.TrimSpace(s)
	}
	return splitted
}

func ConfigFromEnvFile(file string) {
	err := godotenv.Load(file)
	if err != nil {
		panic(err)
	}
	return
}

type Configuration struct {
	Name string `yaml:"name" json:"name"`
	Mode string `yaml:"mode" json:"mode"`
}

var Configur Configuration

func ConfigFromJsonYaml(file string) {

	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Can not open file: %v", err)
		os.Exit(1)
	}

	if strings.Contains(file, ".json") {
		if err = json.Unmarshal(contents, &Configur); err != nil {
			fmt.Printf("Invalid json: %v", err)
			os.Exit(1)
		}
	}

	if strings.Contains(file, ".yaml") {
		if err = yaml.Unmarshal(contents, &Configur); err != nil {
			fmt.Printf("Invalid json: %v", err)
			os.Exit(1)
		}
	}

}
