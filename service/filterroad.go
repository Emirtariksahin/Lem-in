package main

// Filters roads and removes conflicting rooms
func FilterRoads(roads [][]string, antCount int) [][]string {
	var suitableRoads [][]string // Slice to store filtered roads

	// Helper function to check if two roads have overlapping intermediate rooms
	roomsOverlap := func(road1, road2 []string) bool {
		rooms := make(map[string]bool)                 // Map to hold rooms of the first road
		for _, room := range road1[1 : len(road1)-1] { // Exclude start and end rooms
			rooms[room] = true // Add room to the map
		}
		for _, room := range road2[1 : len(road2)-1] { // Check rooms of the second road
			if rooms[room] { // If room exists in the first road, there is an overlap
				return true
			}
		}
		return false // No overlap
	}

	// Try all combinations to find the maximum number of non-overlapping roads
	var tryCombinations func([][]string, int, []int)
	var bestSelection []int // Slice to hold the best selections
	maxRoads := 0           // Maximum number of non-overlapping roads found so far

	// Helper function to try all combinations
	tryCombinations = func(roads [][]string, index int, selected []int) {
		// If the number of selected roads is the most so far, update the best selection
		if len(selected) > maxRoads {
			maxRoads = len(selected)
			bestSelection = make([]int, len(selected))
			copy(bestSelection, selected) // Copy the selections
		}

		// Try remaining roads
		for i := index; i < len(roads); i++ {
			overlaps := false // Flag to check overlap
			for _, sel := range selected {
				// If the new road overlaps with any selected roads, set the flag and break
				if roomsOverlap(roads[sel], roads[i]) {
					overlaps = true
					break
				}
			}
			// If no overlap, add the new road to the selected and try combinations
			if !overlaps {
				selected = append(selected, i)
				tryCombinations(roads, i+1, selected)
				selected = selected[:len(selected)-1] // Backtracking
			}
		}
	}

	// Start combinations
	tryCombinations(roads, 0, []int{})

	// Add suitable roads based on the best selection
	for _, index := range bestSelection {
		suitableRoads = append(suitableRoads, roads[index])
		// If suitable roads reach the ant count, break the loop
		if len(suitableRoads) == antCount {
			break
		}
	}

	return suitableRoads // Return filtered roads
}
