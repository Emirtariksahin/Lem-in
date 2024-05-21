package main

type Node struct {
	Name        string  // Düğümün adı
	Coordinates [2]int  // Düğümün koordinatları
	Edges       []*Edge // Düğüme bağlı kenarlar
}

type Edge struct {
	Start  *Node // Kenarın başlangıç düğümü
	End    *Node // Kenarın bitiş düğümü
	Weight int   // Kenarın ağırlığı (mesafe vb.)
}

type Graph struct {
	Edges []*Edge          // Grafın kenarları
	Nodes map[string]*Node // Grafın düğümleri
}
