package main

import (
	"fmt"
	"strings"
)

// Create a bidirectional graph
func createGraph(kordinatlar map[string][2]int, baglantilar []string) *Graph {
	nodes := make(map[string]*Node)
	var edges []*Edge

	// Create nodes
	for name, coords := range kordinatlar {
		node := &Node{Name: name, Coordinates: coords}
		nodes[name] = node
	}

	// Create edges (bidirectionally)
	for _, conn := range baglantilar {
		parts := strings.Split(conn, "-")
		startNode := nodes[parts[0]]
		endNode := nodes[parts[1]]

		if startNode != nil && endNode != nil {
			// Create edge in both directions
			edge1 := &Edge{Start: startNode, End: endNode, Weight: 1}
			edge2 := &Edge{Start: endNode, End: startNode, Weight: 1}

			startNode.Edges = append(startNode.Edges, edge1)
			endNode.Edges = append(endNode.Edges, edge2)
			edges = append(edges, edge1, edge2)
		}
	}

	// Create the graph
	graph := &Graph{}
	for _, node := range nodes {
		graph.Nodes = append(graph.Nodes, node)
	}
	graph.Edges = edges

	return graph
}

// Print the graph
func (g *Graph) String() string {
	var nodesStr string
	for _, node := range g.Nodes {
		nodesStr += fmt.Sprintf("[%s %v] ", node.Name, node.Coordinates)
	}

	var edgesStr string
	for _, edge := range g.Edges {
		edgesStr += fmt.Sprintf("[%s-%s] ", edge.Start.Name, edge.End.Name)
	}

	return fmt.Sprintf("Nodes: %s\nEdges: %s", nodesStr, edgesStr)
}
