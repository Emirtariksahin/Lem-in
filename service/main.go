package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . istenilen.txt")
		os.Exit(1)
	}

	sentences, err := readInputFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	kordinatlar, err := parseCoordinates(sentences)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	start, end, err := parseStartEndCoordinates(sentences)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	antsayisistring := sentences[0]
	antsayisi, err := strconv.Atoi(antsayisistring)
	if err != nil {
		fmt.Println("Error:", err)

	}
	baglantilar := baglantilar(sentences)
	graph := createGraph(kordinatlar, baglantilar)
	startNode := graph.FindNodeByName(start)
	endNode := graph.FindNodeByName(end)
	fmt.Println("karinca sayisi:", antsayisi)
	fmt.Println("Koordinatlar:", kordinatlar)
	fmt.Println("Başlangiç Koordinati bul:", start)
	fmt.Println("Bitiş Koordinati bul :", end)
	fmt.Println("Baglantilar:", baglantilar)
	fmt.Println("Graf oluşturuldu:", graph)
	// BFS kullanarak tüm doğru yolları bulma
	allPaths := graph.FindAllPathsBFS(startNode, endNode)
	var startsizallpath [][]*Node // start olmadan kısa yollar

	if len(allPaths) == 0 {
		fmt.Println("Başlangıç ve bitiş noktası arasında yol bulunamadı.")
	} else {

		for _, path := range allPaths {

			var currentPath []*Node // Yeni bir yol için geçici bir slice oluştur
			for _, node := range path {
				if startNode != node {
					currentPath = append(currentPath, node)
				}
			}
			// Eğer currentPath boş değilse, startsizallpath'a ekleyin
			if len(currentPath) > 0 {
				startsizallpath = append(startsizallpath, currentPath)
			}
		}
	}
	fmt.Println("Tüm kısa yollar:")
	stringPaths := [][]string{}
	for i, path := range startsizallpath {
		var stringPath []string
		fmt.Printf("Path %d: ", i)
		for _, node := range path {
			fmt.Printf("%s ", node.Name)
			stringPath = append(stringPath, node.Name)
		}
		fmt.Println()
		stringPaths = append(stringPaths, stringPath)
	}
	fmt.Println("String representation of all paths:")
	for i, path := range stringPaths {
		fmt.Printf("Path %d: %v\n", i, path)
	}
	counts := make([]int, len(stringPaths))

	for i := 0; i < len(stringPaths); i++ {
		count := 0
		// Determine the minimum length among all paths
		minPathLength := len(stringPaths[i])
		for _, path := range stringPaths {
			if len(path) < minPathLength {
				minPathLength = len(path)
			}
		}
		// Iterate up to the minimum path length
		for j := 0; j < minPathLength; j++ {
			for k := 0; k < len(stringPaths); k++ {
				if i != k && j < len(stringPaths[k]) && stringPaths[i][j] == stringPaths[k][j] {
					count++
					break
				}
			}
		}
		counts[i] = count
	}
	// Select paths with the minimum count per starting index
	selectedPaths := make([][]string, 0)
	startingIndicesSeen := make(map[string]int)
	for i, path := range stringPaths {
		startIndex := path[0]
		if prevCountIndex, exists := startingIndicesSeen[startIndex]; exists {
			if counts[i] < counts[prevCountIndex] {
				selectedPaths[prevCountIndex] = path[:len(path)-1] // Remove the last index
				startingIndicesSeen[startIndex] = i
			}
		} else {
			selectedPaths = append(selectedPaths, path[:len(path)-1]) // Remove the last index
			startingIndicesSeen[startIndex] = i
		}
	}

	fmt.Println("\nSelected Paths (Min Count per Starting Index) without Last Index:")
	for i, path := range selectedPaths {
		fmt.Printf("Path %d: %v\n", i, path)
	}

	// --- Filtering for Unique Paths (One per Ending Node, Including Empty Paths) ---
	uniqueEndNodePaths := make(map[string][]string)
	for _, path := range selectedPaths {
		endingNode := ""
		if len(path) > 0 {
			endingNode = path[len(path)-1]
		}
		if _, exists := uniqueEndNodePaths[endingNode]; !exists {
			uniqueEndNodePaths[endingNode] = path
		}
	}

	// Convert the map back into a slice of paths
	uniquePaths := [][]string{}
	for _, path := range uniqueEndNodePaths {
		uniquePaths = append(uniquePaths, path)
	}

	// --- Output ---
	fmt.Println("\nFinal Unique Paths (One per Ending Node, Including Empty Paths):")
	for i, path := range uniquePaths {
		fmt.Printf("Path %d: %v\n", i, path)
	}
	finalPaths := [][]string{}
	for _, path := range uniqueEndNodePaths {
		if len(path) == 0 { // If the path is empty, directly append the ending node
			path = append(path, end)
		} else {
			path = append(path, end) // Append the ending node to non-empty paths
		}
		path = append([]string{start}, path...)
		finalPaths = append(finalPaths, path)
	}

	// --- Output ---
	fmt.Println("\nFinal Unique Paths (With Ending Node Appended):")
	for i, path := range finalPaths {
		fmt.Printf("Path %d: %v\n", i, path)
	}
	// Convert finalPaths to slices of *Node
	var finalNodePaths [][]*Node
	for _, path := range finalPaths {
		var nodePath []*Node
		for _, nodeName := range path {
			nodePath = append(nodePath, graph.FindNodeByName(nodeName))
		}
		finalNodePaths = append(finalNodePaths, nodePath)
	}

	// --- Output ---
	fmt.Println("\nFinal Unique Paths (With Ending Node Appended) as Nodes:")
	for i, path := range finalNodePaths {
		fmt.Printf("Path %d: ", i)
		for _, node := range path {
			fmt.Printf("%s ", node.Name)
		}
		fmt.Println()
	}
	println()
	SimulateAnts(graph, antsayisi, startNode, endNode, finalNodePaths)
}

func (graph *Graph) FindNodeByName(name string) *Node {
	for _, node := range graph.Nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}
func SimulateAnts(graph *Graph, ants int, start, end *Node, finalNodePaths [][]*Node) {
	allPaths := finalNodePaths
	if len(allPaths) == 0 {
		fmt.Println("Başlangıç ve bitiş noktası arasında yol bulunamadı.")
		return
	}

	antPaths := make([][]*Node, ants)
	for i := 0; i < ants; i++ {
		antPaths[i] = allPaths[i%len(allPaths)]
	}

	// Track the progress of each ant
	antPositions := make([]int, ants)
	nodeOccupied := make(map[*Node]bool)

	// Initialize ant positions to skip the start node in output
	for i := 0; i < ants; i++ {
		antPositions[i] = 1
	}

	// Simulate until all ants have reached the end
	for {
		allAntsFinished := true
		moveMade := false
		roundOutput := ""

		for i := 0; i < ants; i++ {
			if antPositions[i] < len(antPaths[i]) {
				nextNode := antPaths[i][antPositions[i]]

				if antPositions[i] > 1 && antPositions[i] - 1 < len(antPaths[i]) {
					nodeOccupied[antPaths[i][antPositions[i]-1]] = false
				}

				if !nodeOccupied[nextNode] || nextNode == antPaths[i][len(antPaths[i])-1] {
					roundOutput += fmt.Sprintf("L%d-%s ", i+1, nextNode.Name)
					nodeOccupied[nextNode] = true
					antPositions[i]++ // Move the ant to the next node
					moveMade = true

					if antPositions[i] < len(antPaths[i]) {
						allAntsFinished = false // At least one ant is still moving
					}
				} else {
					allAntsFinished = false
				}
			}
		}

		if moveMade {
			fmt.Println(roundOutput)
		}

		if allAntsFinished {
			break // All ants have finished their paths
		}
	}
}