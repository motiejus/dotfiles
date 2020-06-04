package main

import (
	"fmt"
	"os"

	"github.com/motiejus/dotfiles/joplin2site/internal/cli"
)

func main() {
	a := &cli.App{
		Args:   os.Args[1:],
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := a.Run(); err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
