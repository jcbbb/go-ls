package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var color = "\x1b[%dm%s\x1b[0m"

func main() {
	dir := "./"
	args := os.Args[1:]
	flags := make(map[string]bool)

	for _, v := range args {
		if !strings.HasPrefix(v, "-") {
			dir = v
		} else {
			flags[v] = true
		}
	}

	entries, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") && !flags["-a"] {
			continue
		}

		if e.IsDir() {
			fmt.Println(fmt.Sprintf(color, 34, e.Name()))
		} else {
			fmt.Println(e.Name())
		}
	}
}
