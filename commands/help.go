package commands

import (
	"fmt"

	"github.com/longdivision/goggles/opts"
)

func Help(opts *opts.Options) {
	fmt.Println("Goggles: A command-line search tool for Go packages.")
	fmt.Println()
	fmt.Println("Usage: goggles <query>")
}
