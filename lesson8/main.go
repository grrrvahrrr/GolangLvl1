package main

import (
	"GolangLvl1/lesson8/config"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flagName = flag.String("name", "", "names of people to greet")
	flagMode = flag.String("mode", "", "hello or goodbye")
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

	flag.Parse()

	config.ConfigFromFile("config.env")

	mode := os.Getenv("MODE")
	name := os.Getenv("NAME")
	if *flagMode == "" {
		*flagMode = mode
	}
	if *flagName == "" {
		*flagName = name
	}

	if strings.Contains("hello", *flagMode) {
		Greeter(config.ParseNames(*flagName))
	}

	if strings.Contains("gb", *flagMode) {
		Goodbyer(config.ParseNames(*flagName))
	}
}
