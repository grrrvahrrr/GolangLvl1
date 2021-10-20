package config

import (
	"strings"

	"github.com/joho/godotenv"
)

func ParseNames(raw string) []string {
	splitted := strings.Split(raw, ",")
	for i, s := range splitted {
		splitted[i] = strings.TrimSpace(s)
	}
	return splitted
}

func ConfigFromFile(file string) {
	err := godotenv.Load(file)
	if err != nil {
		panic(err)
	}
	return
}
