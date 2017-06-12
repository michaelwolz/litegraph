// Author: Michael Wolz

package litegraph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Graph -- Graph type with adjacency matrix
type Graph struct {
	vertices  []int
	adjMatrix [][]uint8
}

// Init -- Initializes a graph with v vertices
func (g *Graph) Init(v int) {
	// add vertices
	for i := 0; i < v; i++ {
		g.AddVertex(i)
	}

	// initialize adjacency matrix
	g.adjMatrix = make([][]uint8, v-1)
	for i := range g.adjMatrix {
		g.adjMatrix[i] = make([]uint8, v-1)
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(ID int) {
	g.vertices = append(g.vertices, ID)
}

// AddEdge adds an edge from v1 to v2 to the graph (in the adjacency matrix)
func (g *Graph) AddEdge(v1, v2 int) {
	v1, v2 = minMax(v1, v2)
	g.adjMatrix[v1][v2] = 1
}

// PrintAdjMatrix prints out the adjacency matrix of the graph
// TODO: Fix visualization, atm first row and last coloum are missing
func (g *Graph) PrintAdjMatrix() {
	fmt.Print("\n##### GRAPH ADJACENCY MATRIX #####\n\n")
	for i := 0; i < len(g.vertices)-1; i++ {
		fmt.Printf("(%d) %v\n", i+1, g.adjMatrix[i])
	}
	fmt.Print("\n##################################\n\n")
}

// GenerateJSONGraph outputs the graph to a file in JSON-Format
func (g *Graph) GenerateJSONGraph(path string) {
	json, _ := json.Marshal(g)

	f, err := os.Create(path + "graph.json")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(json))
	check(err)
	w.Flush()

	fmt.Print("JSON-Data written to file: ./graph.json\n\n")
}

// MarshalJSON -- Needed to define this function, because json.Marschal
// messes up uint8 values
func (g *Graph) MarshalJSON() ([]byte, error) {
	var array string
	if g.adjMatrix == nil {
		array = "null"
	} else {
		array = strings.Join(strings.Fields(fmt.Sprintf("%d", g.adjMatrix)), ",")
	}
	jsonResult := fmt.Sprintf(`{"adjMatrix":%s}`, array)
	return []byte(jsonResult), nil
}

//helper functions

//minMax is needed to build/read the adjacency matrix of the graph
func minMax(v1, v2 int) (int, int) {
	// v1 - 1, because we don't have a first row! (lower triangular matrix)
	if v1 > v2 {
		return v1 - 1, v2
	}
	return v2 - 1, v1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
