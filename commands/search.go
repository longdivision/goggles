package commands

import (
	"fmt"
	"strconv"

	"github.com/longdivision/goggles/network"
	"github.com/longdivision/goggles/opts"
)

func Search(opts *opts.Options) {
	fmt.Println("Searching for", opts.Query, "...")

	searchResults := network.PerformGitHubSearch(opts.Query)
	lines := make([]string, 0)
	lines = append(lines, "Results:")

	for i, searchResult := range searchResults.Results {
		lines = append(lines, strconv.Itoa(i)+". "+searchResult.Name+" ("+searchResult.PackagePath+")")
		lines = append(lines, "   "+searchResult.Description)
		lines = append(lines, "")
	}

	lines = lines[:len(lines)-1]

	for _, line := range lines {
		fmt.Println(line)
	}
}
