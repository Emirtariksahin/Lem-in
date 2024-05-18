package main

import "fmt"

func SimulateAnts(graph *Graph, ants int, start, end *Node, allPaths [][]*Node) {
	if len(allPaths) == 0 {
		fmt.Println("Başlangıç ve bitiş noktası arasında yol bulunamadı.")
		return
	}

	antPaths := make([][]*Node, ants)
	for i := 0; i < ants; i++ {
		antPaths[i] = allPaths[i%len(allPaths)]
	}

	// Calculate maximum path length
	maxPathLength := 0
	for _, path := range allPaths {
		if len(path) > maxPathLength {
			maxPathLength = len(path)
		}
	}

	// Track the progress of each ant
	antPositions := make([]int, ants)
	nodeOccupied := make(map[*Node]bool)
	antSteps := make([]int, ants) // Track the number of steps each ant has taken

	// Initialize ant positions and steps (skip the start node in output)
	for i := 0; i < ants; i++ {
		antPositions[i] = 1
		antSteps[i] = 1
	}

	round := 1 // Track the current round

	// Determine the number of connections from the start node
	startNodeConnections := len(start.Edges)

	// Simulate until all ants have reached the end
	for {
		allAntsFinished := true
		roundOutput := fmt.Sprintf("Round %d: ", round)

		// Limit the number of ants moving from the start node each round
		antsMovingFromStart := 0

		for i := 0; i < ants; i++ {
			// Check if the ant has reached the end node
			if antPositions[i] >= len(antPaths[i]) {
				continue // Skip this ant if it's finished
			}

			// Check if the ant should move in this round based on path length
			if antSteps[i] < maxPathLength {
				nextNode := antPaths[i][antPositions[i]]

				if antPositions[i] > 1 && antPositions[i]-1 < len(antPaths[i]) {
					nodeOccupied[antPaths[i][antPositions[i]-1]] = false
				}

				// Limit ants moving from the start node based on its connections
				if antPositions[i] == 1 {
					if antsMovingFromStart >= startNodeConnections {
						continue // Skip this ant if the limit is reached
					}
					antsMovingFromStart++
				}

				if !nodeOccupied[nextNode] || nextNode == antPaths[i][len(antPaths[i])-1] {
					roundOutput += fmt.Sprintf("L%d-%s ", i+1, nextNode.Name)
					nodeOccupied[nextNode] = true
					antPositions[i]++ // Move the ant to the next node
					antSteps[i]++     // Increment the ant's step counter
				}

				if antPositions[i] < len(antPaths[i]) {
					allAntsFinished = false // At least one ant is still moving
				}
			} else {
				allAntsFinished = false // This ant is waiting for others to catch up
			}
		}

		fmt.Println(roundOutput)
		round++

		if allAntsFinished {
			break // All ants have finished their paths
		}
	}
}
