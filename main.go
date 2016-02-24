package main

import (
	"fmt"
	"os"

	"github.com/arvinkulagin/configurator/fill"
	"github.com/arvinkulagin/configurator/get"
)

// Commands
const (
	GET     = "get"
	FILL    = "fill"
	INSTALL = "install"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case GET:
		get.Run()
	case FILL:
		fill.Run()
	default:
		fmt.Println("cfg: unknown command", os.Args[1])
	}
}

func printHelp() {
	fmt.Println("Help")
}
