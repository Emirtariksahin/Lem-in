package main

func containsNode(path []*Node, node *Node) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func (graph *Graph) FindAllPathsBFS(startNode *Node, endNode *Node) [][]*Node {
	var allPaths [][]*Node
	queue := [][]*Node{{startNode}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == endNode {
			allPaths = append(allPaths, path)
			// Eğer son düğüme ulaştıysak, daha fazla genişlemeyi durdur
		}

		for _, edge := range node.Edges {
			if !containsNode(path, edge.End) {
				// Yeni bir dilim oluştur ve mevcut yolu kopyala
				newPath := make([]*Node, len(path))
				copy(newPath, path)
				// Yeni düğümü yeni yola ekle
				newPath = append(newPath, edge.End)
				queue = append(queue, newPath)
			}
		}
	}

	return allPaths
}
