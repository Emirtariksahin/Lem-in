package main

import (
	"strconv"
	"strings"
)

//
// parseCoordinates fonksiyonu, cümlelerden koordinatları çıkarır ve bir harita olarak döndürür
func parseCoordinates(sentences []string) (map[string][2]int, error) {
	coordinates := make(map[string][2]int) // Koordinatları saklayacak harita oluşturulur
	for _, line := range sentences {
		// Satır "-" içermiyorsa, boş değilse ve özel işaretler ##start ve ##end değilse çalışır
		if !strings.Contains(line, "-") && line != "" && line != "##start" && line != "##end" {
			parts := strings.Fields(line) // Satırı boşluklara göre parçalar
			if len(parts) >= 3 {
				key := parts[0]                  // Düğümün adı
				x, err := strconv.Atoi(parts[1]) // X koordinatını dönüştürür
				if err != nil {
					return nil, err // Hata varsa hata döndürür
				}
				y, err := strconv.Atoi(parts[2]) // Y koordinatını dönüştürür
				if err != nil {
					return nil, err // Hata varsa hata döndürür
				}
				coordinates[key] = [2]int{x, y} // Koordinatları haritaya ekler
			}
		}
	}
	return coordinates, nil // Koordinatları ve bir hata olmadığını döndürür
}

// baglantilar fonksiyonu, verilen cümleler içinde bağlantıları (kenarları) ayıklar ve bir dilim olarak döndürür.
func baglantilar(sentences []string) []string {
	var baglanitlar []string // Bağlantıları tutacak bir dilim oluşturulur
	for _, a := range sentences {
		if strings.Contains(a, "-") { // Satır içinde "-" içeriyorsa, bir bağlantı olduğunu varsayar
			baglanitlar = append(baglanitlar, a) // Bağlantıyı dilime ekler
		}
	}
	return baglanitlar // Bağlantıları döndürür
}

// parseStartEndCoordinates fonksiyonu, başlangıç ve bitiş koordinatlarını bulur ve döndürür
func parseStartEndCoordinates(sentences []string) (string, string, error) {
	var startCoord, endCoord string // Başlangıç ve bitiş koordinatlarını tutar

	for i, line := range sentences {
		// Satır ##start ise, bir sonraki satırdaki düğümü başlangıç koordinatı olarak alır
		if line == "##start" {
			startParts := strings.Fields(sentences[i+1]) // Başlangıç koordinatlarını ayırır
			startCoord = startParts[0]                   // Başlangıç koordinatını kaydeder
		} else if line == "##end" {
			endParts := strings.Fields(sentences[i+1]) // Bitiş koordinatlarını ayırır
			endCoord = endParts[0]                     // Bitiş koordinatını kaydeder
		}
	}

	return startCoord, endCoord, nil // Başlangıç ve bitiş koordinatlarını ve bir hata olmadığını döndürür
}
