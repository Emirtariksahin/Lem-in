package main
//
import (
	"fmt"
	"os"
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

	// Karınca sayısını al
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

	// Bilgileri yazdır
	fmt.Println("Karınca sayısı:", antsayisi)
	fmt.Println("Koordinatlar:", kordinatlar)
	fmt.Println("Başlangıç Koordinatı:", start)
	fmt.Println("Bitiş Koordinatı:", end)
	fmt.Println("Bağlantılar:", baglantilar)
	fmt.Println("Graf oluşturuldu:", graph)

	// Tüm yolları bul
	allPaths := graph.FindAllPathsBFS(startNode, endNode)
	// Başlangıç noktası olmadan yolları bul
	startsizallpath := findPathsWithoutStart(allPaths, startNode)

	// Yolları yazdır
	printPaths(startsizallpath)
	// Yolları string formatına çevir
	stringPaths := convertPathsToString(startsizallpath)

	// Yol sayılarını hesapla
	counts := calculatePathCounts(stringPaths)
	// Minimum sayıda yolları seç
	selectedPaths := selectMinCountPaths(stringPaths, counts)

	// Seçilen yolları yazdır
	fmt.Println("\nSeçilen Yollar (Her Başlangıç İndeksi İçin Minimum Sayı):")
	printStringPaths(selectedPaths)

	// Benzersiz yolları filtrele
	uniquePaths := filterUniquePaths(selectedPaths)

	// Benzersiz yolları yazdır
	fmt.Println("\nBenzersiz Yollar (Her Bitiş Düğümü İçin, Boş Yollar Dahil):")
	printStringPaths(uniquePaths)

	// Başlangıç ve bitiş düğümlerini yollara ekle
	finalPaths := appendStartEndToPaths(uniquePaths, start, end)

	// Bitiş düğümü eklenmiş benzersiz yolları yazdır
	fmt.Println("\nBitiş Düğümü Eklenmiş Benzersiz Yollar:")
	printStringPaths(finalPaths)

	// Düğümler olarak bitiş düğümü eklenmiş benzersiz yolları yazdır
	finalNodePaths := convertToNodePaths(finalPaths, graph)

	//stringi düğümlere dönüştürdüğün değerleri yazdır
	fmt.Println("\nDüğümler Olarak Bitiş Düğümü Eklenmiş Benzersiz Yollar:")
	printNodePaths(finalNodePaths)

	//bir boşluk bırak
	println()

	//Karıncaları Hareket Ettir
	SimulateAnts(graph, antsayisi, startNode, endNode, finalNodePaths)
	println()
	// Zaman ölçümü bitir
	elapsed := time.Since(startTime)
	fmt.Printf("Kodun çalışması %.8f saniye sürdü.\n", elapsed.Seconds())
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

// Başlangıç noktası olmadan yolları bulma fonksiyonu
func findPathsWithoutStart(allPaths [][]*Node, startNode *Node) [][]*Node {
	var startsizallpath [][]*Node
	for _, path := range allPaths {
		var currentPath []*Node
		for _, node := range path {
			if startNode != node {
				currentPath = append(currentPath, node)
			}
		}
		if len(currentPath) > 0 {
			startsizallpath = append(startsizallpath, currentPath)
		}
	}
	return startsizallpath
}

// Yolları yazdıran fonksiyon
func printPaths(paths [][]*Node) {
	fmt.Println("Tüm kısa yollar:")
	for i, path := range paths {
		fmt.Printf("Path %d: ", i)
		for _, node := range path {
			fmt.Printf("%s ", node.Name)
		}
		fmt.Println()
	}
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

// Yol sayılarını yani countu hesaplayan fonksiyon
func calculatePathCounts(paths [][]string) []int {
	counts := make([]int, len(paths))
	for i := 0; i < len(paths); i++ {
		count := 0
		minPathLength := len(paths[i])
		for _, path := range paths {
			if len(path) < minPathLength {
				minPathLength = len(path)
			}
		}
		for j := 0; j < minPathLength; j++ {
			for k := 0; k < len(paths); k++ {
				if i != k && j < len(paths[k]) && paths[i][j] == paths[k][j] {
					count++
					break
				}
			}
		}
		counts[i] = count
	}
	return counts
}

// Minimum sayıda yolları seçen fonksiyon
func selectMinCountPaths(paths [][]string, counts []int) [][]string {
	selectedPaths := make([][]string, 0)
	startingIndicesSeen := make(map[string]int)
	for i, path := range paths {
		startIndex := path[0]
		if prevCountIndex, exists := startingIndicesSeen[startIndex]; exists {
			if counts[i] < counts[prevCountIndex] {
				selectedPaths[prevCountIndex] = path[:len(path)-1]
				startingIndicesSeen[startIndex] = i
			}
		} else {
			selectedPaths = append(selectedPaths, path[:len(path)-1])
			startingIndicesSeen[startIndex] = i
		}
	}
	return selectedPaths
}

// Yolları yazdıran fonksiyon (string)
func printStringPaths(paths [][]string) {
	for i, path := range paths {
		fmt.Printf("Path %d: %v\n", i, path)
	}
}

// Benzersiz yolları filtreleyen fonksiyon
func filterUniquePaths(paths [][]string) [][]string {
	uniqueEndNodePaths := make(map[string][]string)
	for _, path := range paths {
		endingNode := ""
		if len(path) > 0 {
			endingNode = path[len(path)-1]
		}
		if _, exists := uniqueEndNodePaths[endingNode]; !exists {
			uniqueEndNodePaths[endingNode] = path
		}
	}
	uniquePaths := [][]string{}
	for _, path := range uniqueEndNodePaths {
		uniquePaths = append(uniquePaths, path)
	}
	return uniquePaths
}

// Başlangıç ve bitiş düğümlerini yollara ekleme fonksiyonu
func appendStartEndToPaths(paths [][]string, start, end string) [][]string {
	finalPaths := [][]string{}
	for _, path := range paths {
		if len(path) == 0 {
			path = append(path, end)
		} else {
			path = append(path, end)
		}
		path = append([]string{start}, path...)
		finalPaths = append(finalPaths, path)
	}
	return finalPaths
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

// Düğümleri yazdıran fonksiyon
func printNodePaths(paths [][]*Node) {
	for i, path := range paths {
		fmt.Printf("Path %d: ", i)
		for _, node := range path {
			fmt.Printf("%s ", node.Name)
		}
		fmt.Println()
	}
}
