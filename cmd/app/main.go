package main

import (
	"os"
)

func main() {
	var err error

	// err := commandline.Execute()
	if err != nil {
		os.Exit(1)
	}
}
