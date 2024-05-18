package main

import (
	"bufio"
	"os"
	"strings"
)

func readInputFile(fileName string) ([]string, error) {
	file, err := os.Open("testler/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func baglantilar(sentences []string) []string {
	var baglanitlar []string
	for _, a := range sentences {
		if strings.Contains(a, "-") {
			baglanitlar = append(baglanitlar, a)
		}
	}
	return baglanitlar
}
