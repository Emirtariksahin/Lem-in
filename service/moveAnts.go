package main

import (
	"fmt"
	"strings"
)

// SimulateAnts fonksiyonu, verilen graf ve yollar ile karıncaların hareketini simüle eder.
func SimulateAnts(graph *Graph, ants int, start, end *Node, allPaths [][]*Node, finalNodePaths []*Node) {
	// Eğer başlangıç ve bitiş noktası arasında yol yoksa uyarı ver ve çık
	if len(allPaths) == 0 {
		fmt.Println("Başlangıç ve bitiş noktası arasında yol bulunamadı.")
		return
	}

	// Karıncaların takip edeceği yolları antPaths dizisine atar
	antPaths := make([][]*Node, ants)
	for i := 0; i < ants; i++ {
		antPaths[i] = allPaths[i%len(allPaths)]
	}
	//son karıncayi kısa yoldan götür
	antPaths[ants-1] = allPaths[0]
	// En uzun yolun uzunluğunu hesapla
	maxPathLength := 0
	for _, path := range allPaths {
		if len(path) > maxPathLength {
			maxPathLength = len(path)
		}
	}

	// Karıncaların pozisyonlarını, adım sayılarını ve düğüm işgal durumlarını tutan diziler oluştur
	antPositions := make([]int, ants)
	nodeOccupied := make(map[*Node]bool)
	antSteps := make([]int, ants)      // Karıncaların attığı adım sayısını takip eder
	finishedAnts := make([]bool, ants) // Bitiş düğümüne ulaşan karıncaları takip eder

	// Başlangıçta tüm karıncaların pozisyonlarını ve adım sayılarını başlat
	for i := 0; i < ants; i++ {
		antPositions[i] = 1
		antSteps[i] = 1
	}

	round := 1 // Mevcut turu takip eder

	// Başlangıç düğümünden çıkan bağlantı sayısını belirler
	startNodeConnections := len(start.Edges)

	// Tüm karıncalar bitiş düğümüne ulaşana kadar simülasyonu döngüde tut
	for {
		allAntsFinished := true // Tüm karıncaların bitip bitmediğini kontrol eder
		roundOutput := fmt.Sprintf("Round %d: ", round)

		// Her turda başlangıç düğümünden hareket eden karıncaların sayısını sınırla
		antsMovingFromStart := 0

		// Her karınca için hareketi kontrol eder
		for i := 0; i < ants; i++ {
			// Eğer karınca bitiş düğümüne ulaştıysa veya bitiş düğümüne ulaşan bir karınca varsa devam et
			if antPositions[i] >= len(antPaths[i]) || finishedAnts[i] {
				continue // Eğer bu karınca bitiş düğümüne ulaşmışsa veya bitiş düğümüne ulaşan bir karınca varsa atla
			}

			// Karınca adımlarını en uzun yola göre kontrol et
			if antSteps[i] < maxPathLength {
				nextNode := antPaths[i][antPositions[i]] // Karıncanın gideceği bir sonraki düğüm

				// Bir önceki düğümün artık boş olduğunu belirt
				if antPositions[i] > 1 && antPositions[i]-1 < len(antPaths[i]) {
					nodeOccupied[antPaths[i][antPositions[i]-1]] = false
				}

				// Başlangıç düğümünden çıkan karıncaların sayısını sınırla
				if antPositions[i] == 1 {
					if antsMovingFromStart >= startNodeConnections {
						continue // Eğer sınır aşıldıysa bu karıncayı atla
					}
					antsMovingFromStart++
				}

				// Eğer düğüm işgal edilmemişse veya bitiş düğümüyse karıncayı hareket ettir
				if !nodeOccupied[nextNode] || nextNode == antPaths[i][len(antPaths[i])-1] {
					roundOutput += fmt.Sprintf("L%d-%s ", i+1, nextNode.Name)
					nodeOccupied[nextNode] = true
					antPositions[i]++ // Karıncayı bir sonraki düğüme taşı
					antSteps[i]++     // Karıncanın adım sayısını arttır

					// Eğer karınca bitiş düğümüne ulaştıysa
					if nextNode == end {
						finishedAnts[i] = true // Karıncayı bitiş düğümüne ulaşmış olarak işaretle
					}
				}

				// Eğer karınca henüz yolunu tamamlamadıysa
				if antPositions[i] < len(antPaths[i]) {
					allAntsFinished = false // En az bir karınca hala hareket ediyor
				}
			} else {
				allAntsFinished = false // Bu karınca diğerlerini bekliyor
			}
		}

		// Her turun çıktısını yazdır
		fmt.Println(strings.TrimSpace(roundOutput))

		// Eğer tüm karıncalar bitiş düğümüne ulaştıysa döngüden çık
		if allAntsFinished {
			break // Tüm karıncalar yollarını tamamladı
		}
		round++
	}

}
