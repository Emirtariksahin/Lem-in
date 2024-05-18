package main

import (
	"bufio"
	"os"
	"strings"
)

// readInputFile fonksiyonu, verilen dosya adından bir dosya okur ve satırları bir dilim olarak döndürür.
func readInputFile(fileName string) ([]string, error) {
	file, err := os.Open("testler/" + fileName) // Dosyayı açar
	if err != nil {
		return nil, err // Hata varsa hata döndürür
	}
	defer file.Close() // Fonksiyon sonunda dosyayı kapatır

	var lines []string                // Okunan satırları tutacak bir dilim oluşturulur
	scanner := bufio.NewScanner(file) // Bir tarama nesnesi oluşturulur
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Satırları dilime ekler
	}
	return lines, scanner.Err() // Satırları ve bir hata olup olmadığını döndürür
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
