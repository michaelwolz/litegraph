// Author: Michael Wolz

package main

import (
	"fmt"
	"os"

	"github.com/michaelwolz/litegraph"
)

func main() {
	var args = os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: shortestPath [path-to-json-file]")
		os.Exit(1)
	}

	g := litegraph.Graph{}
	g.ReadJSONGraph(args[0])
	g.PrintAdjMatrix()
	g.CalculateShortestPaths()
}
