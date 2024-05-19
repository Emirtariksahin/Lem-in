package main
//
// Bir yol içinde bir düğümün olup olmadığını kontrol eden yardımcı fonksiyon
func containsNode(path []*Node, node *Node) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

// BFS kullanarak graf üzerinde tüm yolları bulma fonksiyonu
func (graph *Graph) FindAllPathsBFS(startNode *Node, endNode *Node) [][]*Node {
	var allPaths [][]*Node          // Tüm yolları saklamak için dilim
	queue := [][]*Node{{startNode}} // BFS için kuyruk, başlangıç düğümü ile başlar

	// Kuyruk boşalana kadar döngü
	for len(queue) > 0 {
		path := queue[0]          // Kuyruğun ilk yolunu al
		queue = queue[1:]         // Kuyruğun geri kalanını güncelle
		node := path[len(path)-1] // Mevcut yolun son düğümünü al

		// Eğer son düğüme ulaştıysak
		if node == endNode {
			allPaths = append(allPaths, path) // Yolu allPaths dilimine ekle
			// Daha fazla genişlemeyi durdur
		}

		// Mevcut düğümün tüm komşularını kontrol et
		for _, edge := range node.Edges {
			if !containsNode(path, edge.End) { // Eğer komşu düğüm mevcut yolda yoksa
				// Yeni bir dilim oluştur ve mevcut yolu kopyala
				newPath := make([]*Node, len(path))
				copy(newPath, path)
				// Yeni düğümü yeni yola ekle
				newPath = append(newPath, edge.End)
				queue = append(queue, newPath) // Yeni yolu kuyrukta sıraya ekle
			}
		}
	}

	return allPaths // Tüm bulunan yolları döndür
}
