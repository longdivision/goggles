package main

import (
	"os"

	"github.com/longdivision/goggles/commands"
	"github.com/longdivision/goggles/opts"
)

func main() {
	commandToHandler := make(map[string]commands.Command)
	commandToHandler["search"] = commands.Search
	commandToHandler["help"] = commands.Help

	parsedOpts := opts.ParseArgs(os.Args)
	command := commandToHandler[parsedOpts.Command]

	command(parsedOpts)
}
