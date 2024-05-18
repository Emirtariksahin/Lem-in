package main

import (
	"strconv"
	"strings"
)

func parseCoordinates(sentences []string) (map[string][2]int, error) {
	coordinates := make(map[string][2]int)
	for _, line := range sentences {
		if !strings.Contains(line, "-") && line != "" && line != "##start" && line != "##end" {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				key := parts[0]
				x, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, err
				}
				y, err := strconv.Atoi(parts[2])
				if err != nil {
					return nil, err
				}
				coordinates[key] = [2]int{x, y}
			}
		}
	}
	return coordinates, nil
}
func parseStartEndCoordinates(sentences []string) (string, string, error) {
	var startCoord, endCoord string

	for i, line := range sentences {
		if line == "##start" {
			startParts := strings.Fields(sentences[i+1])
			startCoord = startParts[0]
		} else if line == "##end" {
			endParts := strings.Fields(sentences[i+1])
			endCoord = endParts[0]
		}
	}

	return startCoord, endCoord, nil
}
