package main

//
import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Ana fonksiyon
func main() {
	// Zaman ölçümü başlat
	startTime := time.Now()
	// Komut satırı argümanlarını kontrol et
	if len(os.Args) != 2 {
		fmt.Println("Kullanım: go run . istenilen.txt")
		os.Exit(1)
	}

	// Girdi dosyasını oku
	sentences, err := readInputFile(os.Args[1])
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	// Koordinatları ayrıştır
	kordinatlar, err := parseCoordinates(sentences)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	// Başlangıç ve bitiş koordinatlarını ayrıştır
	start, end, err := parseStartEndCoordinates(sentences)
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	for _, line := range sentences {
		fmt.Println(string(line))
	}
	/// Karınca sayısını al
	antsayisistring := sentences[0]
	antsayisi, err := strconv.Atoi(antsayisistring)
	if err != nil {
		fmt.Println("Hata:", err)
	}
	//eğer karınca sayısı 0 sa hata mesajı dönder
	if antsayisi == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	// Bağlantıları ayıkla
	baglantilar := baglantilar(sentences)

	// Bağlantıları işleme ,ona göre hata mesajı döndürme
	var once []string  // "-" işaretinden öncekiler
	var sonra []string // "-" işaretinden sonrakiler

	for _, baglanti := range baglantilar {
		// "-" işaretiyle ayır
		parts := strings.Split(baglanti, "-")
		if len(parts) == 2 {
			once = append(once, parts[0])
			sonra = append(sonra, parts[1])
		}
	}

	for i := 0; i < len(once); i++ {

		if string(once[i]) == string(sonra[i]) {
			fmt.Println("ERROR: invalid data format")
			return

		}
	}

	// Graf oluştur
	graph := createGraph(kordinatlar, baglantilar)
	startNode := graph.FindNodeByName(start)
	endNode := graph.FindNodeByName(end)

	// Tüm yolları bul
	allPaths := graph.FindAllPathsBFS(startNode, endNode)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Yolları string formatına çevir
	stringPaths := convertPathsToString(allPaths)

	filtrelenmisyollar := FilterRoads(stringPaths, antsayisi)

	// Düğümler olarak bitiş düğümü eklenmiş benzersiz yolları yazdır
	finalNodePaths := convertToNodePaths(filtrelenmisyollar, graph)
	a := finalNodePaths[0]

	println()
	//Karıncaları Hareket Ettir
	SimulateAnts(graph, antsayisi, startNode, endNode, finalNodePaths, a)

	// Zaman ölçümü bitir
	elapsed := time.Since(startTime)
	println()
	fmt.Printf("This code took %.8f seconds to run.\n", elapsed.Seconds())
}

// Graf düğümlerini isimle bulma fonksiyonu
func (graph *Graph) FindNodeByName(name string) *Node {
	for _, node := range graph.Nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

// Yolları string formatına çeviren fonksiyon
func convertPathsToString(paths [][]*Node) [][]string {
	var stringPaths [][]string
	for _, path := range paths {
		var stringPath []string
		for _, node := range path {
			stringPath = append(stringPath, node.Name)
		}
		stringPaths = append(stringPaths, stringPath)
	}
	return stringPaths
}

// Yolları düğüm olarak dönüştüren fonksiyon
func convertToNodePaths(paths [][]string, graph *Graph) [][]*Node {
	var finalNodePaths [][]*Node
	for _, path := range paths {
		var nodePath []*Node
		for _, nodeName := range path {
			nodePath = append(nodePath, graph.FindNodeByName(nodeName))
		}
		finalNodePaths = append(finalNodePaths, nodePath)
	}
	return finalNodePaths
}
