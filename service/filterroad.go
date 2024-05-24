package main

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
