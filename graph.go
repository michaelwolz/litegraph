// Author: Michael Wolz

package litegraph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// Graph -- Graph type with adjacency matrix
type Graph struct {
	Vertices  int
	AdjMatrix []uint8
}

// Init -- Initializes a graph with v vertices
func (g *Graph) Init(v int) {
	g.Vertices = v

	// initialize adjacency matrix
	g.AdjMatrix = make([]uint8, v*(v-1)/2)
}

// AddEdge adds an edge from v1 to v2 to the graph
func (g *Graph) AddEdge(v1, v2 int) {
	// if v1 != v2 { // we don't allow loops at this point
	// 	g.AdjMatrix[v1+v2-1] = 1
	// }
	g.AdjMatrix[getIndex(v1, v2)] = 1
}

// RemoveEdge removes an edge from v1 to v2 from the graph (NOT IN USE!)
func (g *Graph) RemoveEdge(v1, v2 int) {
	g.AdjMatrix[getIndex(v1, v2)] = 0
}

// AddRandomEdge adds an random edge to the graph
func (g *Graph) AddRandomEdge() {
	v1, v2 := getRandomVertices(g.Vertices)
	index := getIndex(v1, v2)

	if g.AdjMatrix[index] == 0 {
		g.AdjMatrix[index] = 1
	} else {
		g.AddRandomEdge()
	}
}

// RemoveRandomEdge removes an random edge from the graph
func (g *Graph) RemoveRandomEdge() {
	v1, v2 := getRandomVertices(g.Vertices)
	index := getIndex(v1, v2)

	if g.AdjMatrix[index] == 1 {
		g.AdjMatrix[index] = 0
	} else {
		g.RemoveRandomEdge()
	}
}

func getRandomVertices(v int) (int, int) {
	var v1, v2 int

	for v1 == v2 {
		v1 = rand.Intn(v - 1)
		v2 = rand.Intn(v - 1)
	}

	return v1, v2
}

func getIndex(v1, v2 int) int {
	v1, v2 = minMax(v1, v2)
	return v2*(v2-1)/2 + v1
}

// ConnectAll connects all vertices in the grah
func (g *Graph) ConnectAll() {
	for i := range g.AdjMatrix {
		g.AdjMatrix[i] = 1
	}
}

// CalculateShortestPaths calculates the shortest paths between all vertices
func (g *Graph) CalculateShortestPaths() {
	start := time.Now()

	for i := 0; i < g.Vertices; i++ {
		for j := 0; j < i; j++ {
			g.calculateShortestPath(i, j)
		}
	}

	elapsed := time.Since(start)
	log.Printf("Calculation of shortest paths took %s", elapsed)
}

// CalculateShortestPath calculates the shortest path between two given vertices
// using breadth first search
func (g *Graph) calculateShortestPath(v1, v2 int) {
	var queue = []int{}
	isVisited := make([]bool, g.Vertices)
	isVisited[v1] = true
	counter := 0

	fmt.Printf("### Searching for shortes path betweed %d and %d ###\n", v1, v2)

	//add v1 to queue
	queue = append(queue, v1)

BFS:
	for len(queue) > 0 {
		v := queue[0]
		counter++

		for i := 0; i < g.Vertices; i++ {
			if v != i && g.AdjMatrix[getIndex(v, i)] == 1 {
				if i == v2 {
					fmt.Printf("Found after %d hop(s) at vertice %d!\n\n", counter, v)
					break BFS
				} else if !isVisited[i] {
					isVisited[i] = true
					queue = append(queue, i)
				}
			}
		}

		queue = queue[1:]
	}
}

func (g *Graph) areConnected(v1, v2 int) bool {
	return g.AdjMatrix[getIndex(v1, v2)] == 1
}

// PrintAdjMatrix prints out the adjacency matrix of the graph
// TODO: Fix visualization, atm first row and last column are missing
func (g *Graph) PrintAdjMatrix() {
	fmt.Print("\n#### GRAPH ADJACENCY MATRIX ####\n\n")
	for i := 1; i < g.Vertices; i++ {
		fmt.Printf("(%d) ", i)
		for j := 0; j < i; j++ {
			fmt.Printf("%v ", g.AdjMatrix[i*(i-1)/2+j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nArray: %v\n", g.AdjMatrix)
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
}

//helper functions

func minMax(v1, v2 int) (int, int) {
	if v1 < v2 {
		return v1, v2
	}
	return v2, v1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
