package main

import (
	"GolangLvl1/lesson9/config"
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

	fmt.Println(`Do you want to read from json/yaml file? "yes" or "no"`)

	var readfrom string
	if _, err := fmt.Scan(&readfrom); err != nil {
		fmt.Printf("Invalid answer: %v", err)
		os.Exit(1)
	}

	switch readfrom {
	case "yes":
		fmt.Println(`Please input file name`)
		var filename string
		if _, err := fmt.Scan(&filename); err != nil {
			fmt.Printf("Invalid filename: %v", err)
			os.Exit(1)
		}
		config.ConfigFromJsonYaml(filename)

		*flagName = config.Configur.Name
		*flagMode = config.Configur.Mode

	case "no":
		config.ConfigFromEnvFile("config.env")

		mode := os.Getenv("MODE")
		name := os.Getenv("NAME")

		if *flagMode == "" {
			*flagMode = mode
		}
		if *flagName == "" {
			*flagName = name
		}

	default:
		fmt.Println("Invalid answer")
		os.Exit(1)
	}

	if strings.Contains("hello", *flagMode) {
		Greeter(config.ParseNames(*flagName))
	}

	if strings.Contains("gb", *flagMode) {
		Goodbyer(config.ParseNames(*flagName))
	}
}
