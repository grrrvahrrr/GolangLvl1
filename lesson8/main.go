package main

import (
	"GolangLvl1/lesson8/config"
	"flag"
	"fmt"
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
	c.LoadConfig(*flagName, *flagMode)

	if strings.Contains("hello", c.MODE) {
		Greeter(c.ParseNames())
	}

	if strings.Contains("gb", c.MODE) {
		Goodbyer(c.ParseNames())
	}

}
