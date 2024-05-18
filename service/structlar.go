package main

type Node struct {
	Name        string
	Coordinates [2]int
	Occupied    bool
	Visited     bool
	AntID       int
	Distance    float64
	Previous    *Node
	Edges       []*Edge // Düğümün kenarları
	AntCount    int     // Düğümdeki karınca sayısı
}

type Edge struct {
	Start  *Node
	End    *Node
	Weight int // Initial weight (distance)
	// Additional field for dynamic weight (optional for ACO)
	Distance float64
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}
