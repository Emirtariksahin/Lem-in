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

	filtrelenmisyollar := FilterRoad(stringPaths, antsayisi)

	// Düğümler olarak bitiş düğümü eklenmiş benzersiz yolları yazdır
	finalNodePaths := convertToNodePaths(filtrelenmisyollar, graph)
	a := finalNodePaths[0]

	println()
	//Karıncaları Hareket Ettir
	SimulateAnts(graph, antsayisi, startNode, endNode, finalNodePaths, a)

	// Zaman ölçümü bitir
	elapsed := time.Since(startTime)
	println()
	fmt.Printf("Kodun çalışması %.8f saniye sürdü.\n", elapsed.Seconds())
}

// Yolları filtreler ve çakışan odaları çıkarır
func FilterRoad(yollar [][]string, karincaSayisi int) [][]string {
	var uygunYollar [][]string // Filtrelenmiş yolları saklamak için dilim

	// İki yolun ara odalarda çakışıp çakışmadığını kontrol eden yardımcı fonksiyon
	odalarCakisiyorMu := func(yol1, yol2 []string) bool {
		odalar := make(map[string]bool)             // İlk yolun odalarını tutmak için bir harita
		for _, oda := range yol1[1 : len(yol1)-1] { // Başlangıç ve bitiş odalarını hariç tut
			odalar[oda] = true // Haritaya oda ekle
		}
		for _, oda := range yol2[1 : len(yol2)-1] { // İkinci yolun odalarını kontrol et
			if odalar[oda] { // Eğer oda ilk yolda varsa, çakışma var demektir
				return true
			}
		}
		return false // Çakışma yok
	}

	// En fazla sayıda çakışmayan yolu bulmak için tüm kombinasyonları dene
	var kombinasyonlariDeneyerekBul func([][]string, int, []int)
	var enIyiSecim []int // En iyi seçimleri tutmak için dilim
	enFazlaYol := 0      // Şu ana kadar bulunan en fazla çakışmayan yol sayısı

	// Tüm kombinasyonları denemek için yardımcı fonksiyon
	kombinasyonlariDeneyerekBul = func(yollar [][]string, index int, secilenler []int) {
		// Eğer seçilen yolların sayısı şu ana kadarki en fazlaysa, en iyi seçimi güncelle
		if len(secilenler) > enFazlaYol {
			enFazlaYol = len(secilenler)
			enIyiSecim = make([]int, len(secilenler))
			copy(enIyiSecim, secilenler) // Seçimleri kopyala
		}

		// Kalan yolları dene
		for i := index; i < len(yollar); i++ {
			cakisiyor := false // Çakışma durumunu kontrol etmek için bayrak
			for _, sec := range secilenler {
				// Eğer seçilmiş yollarla yeni yol çakışıyorsa, bayrağı ayarla ve döngüyü kır
				if odalarCakisiyorMu(yollar[sec], yollar[i]) {
					cakisiyor = true
					break
				}
			}
			// Eğer çakışma yoksa, yeni yolu seçilenlere ekle ve kombinasyonları dene
			if !cakisiyor {
				secilenler = append(secilenler, i)
				kombinasyonlariDeneyerekBul(yollar, i+1, secilenler)
				secilenler = secilenler[:len(secilenler)-1] // Geriye doğru izleme (backtracking)
			}
		}
	}

	// Kombinasyonları başlat
	kombinasyonlariDeneyerekBul(yollar, 0, []int{})

	// En iyi seçimlere göre uygun yolları ekle
	for _, index := range enIyiSecim {
		uygunYollar = append(uygunYollar, yollar[index])
		// Eğer uygun yollar karınca sayısına ulaştıysa, döngüyü kır
		if len(uygunYollar) == karincaSayisi {
			break
		}
	}

	return uygunYollar // Filtrelenmiş yolları döndür
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
