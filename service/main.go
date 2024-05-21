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

	// Bilgileri yazdır
	fmt.Println("Karınca sayısı:", antsayisi)
	fmt.Println("Koordinatlar:", kordinatlar)
	fmt.Println("Başlangıç Koordinatı:", start)
	fmt.Println("Bitiş Koordinatı:", end)
	fmt.Println("Bağlantılar:", baglantilar)
	fmt.Println("Graf oluşturuldu:", graph)

	// Tüm yolları bul
	allPaths := graph.FindAllPathsBFS(startNode, endNode)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Yolları string formatına çevir
	stringPaths := convertPathsToString(allPaths)

	filtrelenmisyollar := FiltreleYollar(stringPaths, antsayisi)

	// Düğümler olarak bitiş düğümü eklenmiş benzersiz yolları yazdır
	finalNodePaths := convertToNodePaths(filtrelenmisyollar, graph)
	a := finalNodePaths[0]
	//stringi düğümlere dönüştürdüğün değerleri yazdır
	fmt.Println("\nDüğümler Olarak Bitiş Düğümü Eklenmiş Benzersiz Filtrelenmiş Yollar:")
	printNodePaths(finalNodePaths)

	//bir boşluk bırak
	println()

	//Karıncaları Hareket Ettir
	SimulateAnts(graph, antsayisi, startNode, endNode, finalNodePaths, a)
	println()

	// Zaman ölçümü bitir
	elapsed := time.Since(startTime)
	fmt.Printf("Kodun çalışması %.8f saniye sürdü.\n", elapsed.Seconds())
}

// Yolları filtreler ve çakışan odaları çıkarır
func FiltreleYollar(yollar [][]string, karincaSayisi int) [][]string {
	var filtrelenmisYollar [][]string

	// İki yolun ara odalarda çakışıp çakışmadığını kontrol eden yardımcı fonksiyon
	yollarCakisiyor := func(yol1, yol2 []string) bool {
		kume := make(map[string]bool)
		for _, oda := range yol1[1 : len(yol1)-1] { // Başlangıç ve bitişi hariç tut
			kume[oda] = true
		}
		for _, oda := range yol2[1 : len(yol2)-1] {
			if kume[oda] {
				return true
			}
		}
		return false
	}

	// Çakışmayan yol kombinasyonlarını bulmak için tüm kombinasyonları dene
	var kombinasyonlar func([][]string, int, []int)
	var enIyiKombinasyon []int
	maxYol := 0

	kombinasyonlar = func(yollar [][]string, indeks int, secili []int) {
		if len(secili) > maxYol {
			maxYol = len(secili)
			enIyiKombinasyon = make([]int, len(secili))
			copy(enIyiKombinasyon, secili)
		}

		for i := indeks; i < len(yollar); i++ {
			cakisiyor := false
			for _, s := range secili {
				if yollarCakisiyor(yollar[s], yollar[i]) {
					cakisiyor = true
					break
				}
			}
			if !cakisiyor {
				secili = append(secili, i)
				kombinasyonlar(yollar, i+1, secili)
				secili = secili[:len(secili)-1]
			}
		}
	}

	kombinasyonlar(yollar, 0, []int{})

	for _, indeks := range enIyiKombinasyon {
		filtrelenmisYollar = append(filtrelenmisYollar, yollar[indeks])
		if len(filtrelenmisYollar) == karincaSayisi {
			break
		}
	}

	return filtrelenmisYollar
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
