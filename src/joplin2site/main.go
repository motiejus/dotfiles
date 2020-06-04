package main

import (
	"fmt"
	"os"
)

func main() {
	a := &app{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}

	if err := a.run(); err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
