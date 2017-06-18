// Author: Michael Wolz

package litegraph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Graph -- Graph type with adjacency matrix
type Graph struct {
	vertices  []int
	AdjMatrix []uint8
}

// Init -- Initializes a graph with v vertices
func (g *Graph) Init(v int) {
	// add vertices
	for i := 0; i < v; i++ {
		g.AddVertex(i)
	}

	// initialize adjacency matrix
	g.AdjMatrix = make([]uint8, v*(v-1)/2)
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(ID int) {
	g.vertices = append(g.vertices, ID)
}

// AddEdge adds an edge from v1 to v2 to the graph (in the adjacency matrix)
func (g *Graph) AddEdge(v1, v2 int) {
	if v1 != v2 { // we don't allow loops at this point
		g.AdjMatrix[v1+v2-1] = 1
	}
}

// PrintAdjMatrix prints out the adjacency matrix of the graph
// TODO: Fix visualization, atm first row and last coloum are missing
func (g *Graph) PrintAdjMatrix() {
	fmt.Println(g.AdjMatrix)
	fmt.Print("\n#### GRAPH ADJACENCY MATRIX ####\n\n")
	for i := 0; i < len(g.vertices)-1; i++ {
		fmt.Printf("(%d) ", i+1)
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", g.AdjMatrix[i+j])
		}
		fmt.Printf("\n")
	}
	fmt.Print("\n################################\n\n")
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

// ReadJSONGraph reads in graph from JSON-File
func (g *Graph) ReadJSONGraph(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)

	json.Unmarshal(dat, &g)
	fmt.Println(g.AdjMatrix)
}

//
// // MarshalJSON -- Needed to define this function, because json.Marschal
// // messes up uint8 values
// func (g *Graph) MarshalJSON() ([]byte, error) {
// 	var array string
// 	if g.AdjMatrix == nil {
// 		array = "null"
// 	} else {
// 		array = strings.Join(strings.Fields(fmt.Sprintf("%d", g.AdjMatrix)), ",")
// 	}
// 	jsonResult := fmt.Sprintf(`{"AdjMatrix":%s}`, array)
// 	return []byte(jsonResult), nil
// }
//
// // UnmarshalJSON -- Needed to define this function, because json.Marschal
// func (g *Graph) UnmarshalJSON([]byte) {
// 	var array string
// 	if g.AdjMatrix == nil {
// 		array = "null"
// 	} else {
// 		array = strings.Join(strings.Fields(fmt.Sprintf("%d", g.AdjMatrix)), ",")
// 	}
// 	jsonResult := fmt.Sprintf(`{"AdjMatrix":%s}`, array)
// 	return []byte(jsonResult), nil
// }

//helper functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}
