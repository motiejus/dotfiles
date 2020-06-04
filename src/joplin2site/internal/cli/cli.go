package cli

import (
	"fmt"
	"io"

	goflags "github.com/jessevdk/go-flags"
)

type flags struct {
	PositionalArgs struct {
		Dir goflags.Filename `long:"dir" description:"Directory with Joplin notes"`
	} `positional-args:"yes" required:"yes"`
}

type App struct {
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (a *App) Run() error {
	var opts flags
	args, err := goflags.ParseArgs(&opts, a.Args)
	if err != nil {
		return err
	}
	if len(args) != 0 {
		return fmt.Errorf("Got unexpected arguments: %q", args)
	}
	return nil
}
