package main

//
import (
	"strings"
)

// Bidirectional (çift yönlü) bir graf oluşturma fonksiyonu
func createGraph(kordinatlar map[string][2]int, baglantilar []string) *Graph {
	nodes := make(map[string]*Node) // Düğümleri saklayacak harita
	var edges []*Edge               // Kenarları saklayacak dilim

	// Düğümleri oluştur
	for name, coords := range kordinatlar {
		node := &Node{Name: name, Coordinates: coords}
		nodes[name] = node
	}

	// Kenarları oluştur (çift yönlü)
	for _, conn := range baglantilar {
		parts := strings.Split(conn, "-")
		startNode := nodes[parts[0]]
		endNode := nodes[parts[1]]

		// Başlangıç ve bitiş düğümleri varsa
		if startNode != nil && endNode != nil {
			// Her iki yönde kenar oluştur
			edge1 := &Edge{Start: startNode, End: endNode, Weight: 1}
			edge2 := &Edge{Start: endNode, End: startNode, Weight: 1}

			startNode.Edges = append(startNode.Edges, edge1)
			endNode.Edges = append(endNode.Edges, edge2)
			edges = append(edges, edge1, edge2)
		}
	}

	// Grafı oluştur
	graph := &Graph{Nodes: nodes, Edges: edges}
	return graph
}
