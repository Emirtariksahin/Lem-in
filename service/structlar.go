package main
//
// Node, graf yapısındaki düğümleri temsil eder.
type Node struct {
	Name        string  // Düğümün adı
	Coordinates [2]int  // Düğümün koordinatları
	Edges       []*Edge // Düğüme bağlı kenarlar
}

// Edge, graf yapısındaki kenarları temsil eder.
type Edge struct {
	Start  *Node // Kenarın başlangıç düğümü
	End    *Node // Kenarın bitiş düğümü
	Weight int   // Kenarın başlangıç ağırlığı (mesafe vb.)
}

// Graph, graf yapısını temsil eder.
type Graph struct {
	Nodes []*Node // Grafın düğümleri
	Edges []*Edge // Grafın kenarları
}
