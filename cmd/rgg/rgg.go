// Author: Michael Wolz

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/michaelwolz/litegraph"
)

var maxEdges int

func main() {
	var args = os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: rgg vertices edges")
		os.Exit(1)
	}

	v := argParse(args[0])
	e := argParse(args[1])
	maxEdges = v * (v - 1) / 2

	if e > maxEdges {
		fmt.Fprintf(os.Stderr, "error: max amount of edges in a graph with %d vertices is %d. Edge weight must be 1\n", v, maxEdges)
		os.Exit(1)
	}

	if e < v-1 {
		fmt.Fprintf(os.Stderr, "error: min amount of edges in a graph with %d vertices is %d. Edge weight must be 1\n", v, v-1)
		os.Exit(1)
	}

	// seed the pseudo-rand generator
	rand.Seed(time.Now().UnixNano())

	// build graph
	g := buildRandomGraph(v, e)

	// ouput adjList (for debugging only)
	g.PrintAdjMatrix()

	// write graph to JSON-file
	g.GenerateJSONGraph("./")
}

func buildRandomGraph(v, e int) litegraph.Graph {
	var g = litegraph.Graph{}
	g.Init(v)

	if e == maxEdges {
		g.ConnectAll()
	} else {
		if e < (v+maxEdges)/2 {
			distributeEdges(g, v, e)
		} else {
			removeEdges(g, v, e)
		}
	}

	return g
}

func distributeEdges(g litegraph.Graph, v, e int) {
	buildBaseGraph(g, v)

	// add remaining edges to the graph
	remaining := e - v + 1
	for remaining > 0 {
		g.AddRandomEdge()
		remaining--
	}
}

func removeEdges(g litegraph.Graph, v, e int) {
	g.ConnectAll()

	// remove additional edges
	remaining := maxEdges - e
	fmt.Println(remaining)
	for remaining > 0 {
		g.RemoveRandomEdge()
		remaining--
	}
}

func buildBaseGraph(g litegraph.Graph, v int) {
	// connect all vertices with v-1 edges
	var vertexPermutation = rand.Perm(v)

	for i := 0; i < len(vertexPermutation)-1; i++ {
		g.AddEdge(vertexPermutation[i], vertexPermutation[i+1])
	}
}

// ### helper functions ###

// we need this to parse integer values from the terminal
func argParse(arg string) int {
	res, err := strconv.Atoi(arg)
	check(err)
	return res
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
