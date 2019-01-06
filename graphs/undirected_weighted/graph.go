package main

import "fmt"

// Node is a single node in a graph list
type Node struct {
	name  string
	value int
}

func addVertexToGraph(vtx string) {
	if graph[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	graph[vtx] = make([]Node, 0)
}

func addEdgeToGraph(fromVtx, toVtx string, edgeValue int) {
	if graph[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range graph[fromVtx] { // check if edge already exists
		if graph[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}
	if graph[toVtx] == nil { // create new destination vertext if it does not exists
		graph[toVtx] = make([]Node, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	graph[fromVtx] = append(graph[fromVtx], Node{name: toVtx, value: edgeValue})
	graph[toVtx] = append(graph[toVtx], Node{name: fromVtx, value: edgeValue})
	return
}

func removeVertexFromGraph(vtx string) {
	length := len(graph[vtx]) - 1
	for length != -1 {
		removeEdgeFromGraph(vtx, graph[vtx][length].name)
		length--
	}
	delete(graph, vtx)
}

func removeEdgeFromGraph(fromVtx, toVtx string) {
	if graph[fromVtx] == nil || graph[toVtx] == nil {
		fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " does not exist. --")
		return
	}

	for i := range graph[fromVtx] {
		if graph[fromVtx][i].name == toVtx {
			if i == 0 {
				graph[fromVtx] = graph[fromVtx][1:len(graph[fromVtx])]
			} else if i == (len(graph[fromVtx]) - 1) {
				graph[fromVtx] = graph[fromVtx][0:(len(graph[fromVtx]) - 1)]
			} else {
				initial := graph[fromVtx][0:i]
				final := graph[fromVtx][i+1 : len(graph[fromVtx])]
				graph[fromVtx] = append(initial, final...)
			}
			break
		}
	}

	for i := range graph[toVtx] {
		if graph[toVtx][i].name == fromVtx {
			if i == 0 {
				graph[toVtx] = graph[toVtx][1:len(graph[toVtx])]
			} else if i == (len(graph[toVtx]) - 1) {
				graph[toVtx] = graph[toVtx][0:(len(graph[toVtx]) - 1)]
			} else {
				initial := graph[toVtx][0:i]
				final := graph[toVtx][i+1 : len(graph[toVtx])]
				graph[toVtx] = append(initial, final...)
			}
			break
		}
	}
}

var graph map[string][]Node

func init() {
	graph = make(map[string][]Node)
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. REMOVE VERTEX")
		fmt.Println("4. REMOVE AN EDGE")
		fmt.Println("5. DISPLAY USING DFS")
		fmt.Println("6. DISPLAY USING BFS")
		fmt.Println("7. SIMPLE DISPLAY")
		fmt.Println("8. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			addVertex()
			break
		case 2:
			addEdge()
			break
		case 3:
			removeVertex()
			break
		case 4:
			removeEdge()
			break
		case 5:
			result, weights := displayDFS()
			fmt.Println(result)
			fmt.Println(weights)
			break
		case 6:
			result, weights := displayBFS()
			fmt.Println(result)
			fmt.Println(weights)
			break
		case 7:
			simpleDisplay()
			break
		case 8:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	addVertexToGraph(vtxName)
}

func addEdge() {
	var fromVtx, toVtx string
	var edgeValue int
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d", &edgeValue)
	addEdgeToGraph(fromVtx, toVtx, edgeValue)
}

func removeVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	removeVertexFromGraph(vtxName)
}

func removeEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	removeEdgeFromGraph(fromVtx, toVtx)
}

func displayDFS() ([]string, []int) {
	var startVtx string
	fmt.Print("Enter the start vertex name: ")
	fmt.Scanf("%s", &startVtx)
	result, weights := DFS(startVtx)
	return result, weights
}

func displayBFS() ([]string, []int) {
	var startVtx string
	fmt.Print("Enter the start vertex name: ")
	fmt.Scanf("%s", &startVtx)
	result, weights := BFS(startVtx)
	return result, weights
}

func simpleDisplay() {
	fmt.Println("")
	for i := range graph {
		fmt.Print(i, " => ")
		for j := range graph[i] {
			fmt.Print(graph[i][j])
		}
		fmt.Println("")
	}
}
