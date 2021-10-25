//Lesson 9 Homework
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
	var c config.Configuration

	fmt.Println(`Please input config file name or "d" for default`)
	var filename string
	if _, err := fmt.Scan(&filename); err != nil {
		fmt.Printf("Invalid filename: %v", err)
		os.Exit(1)
	}

	if filename != "d" {
		err := c.ConfigFromJsonYaml(filename)
		if err != nil {
			fmt.Printf("Invalid file: %v \n", err)
			os.Exit(1)
		}
	} else if filename == "d" {
		c.LoadConfig(*flagName, *flagMode)
	}

	if strings.Contains("hello", c.MODE) {
		Greeter(c.ParseNames())
	}
	if strings.Contains("gb", c.MODE) {
		Goodbyer(c.ParseNames())
	}
}
