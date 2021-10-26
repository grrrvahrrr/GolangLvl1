package main

import (
	"GolangLvl1/lesson8/config"
	"fmt"
	"strings"
)

func Greeter(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("| %s | %s |\n", "Hello", names[i])
	}
	return
}

func Goodbyer(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("| %s | %s |\n", "Goodbye", names[i])
	}
	return
}

func main() {
	cfg, err := config.Load("config_example.env")
	if err != nil {
		fmt.Printf("Couldn't load config %s", err)
	}

	if strings.Contains("hello", cfg.Mode) {
		Greeter(cfg.ParseNames())
	}

	if strings.Contains("gb", cfg.Mode) {
		Goodbyer(cfg.ParseNames())
	}

}
