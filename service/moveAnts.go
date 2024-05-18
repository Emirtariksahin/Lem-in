package main

import "fmt"

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

				if antPositions[i] > 1 && antPositions[i]-1 < len(antPaths[i]) {
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
