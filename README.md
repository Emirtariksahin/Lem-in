
# Lem-in

Karınca Optimizasyon Algoritması ile Kısa Yol Bulma Uygulaması
Bu uygulama, karınca optimizasyon algoritmasını kullanarak bir graf üzerinde en kısa yolu bulmayı sağlar. Kullanıcı, girdi dosyasında grafın yapısını, düğümlerin koordinatlarını ve başlangıç ile bitiş düğümlerini tanımlar. Program daha sonra karıncaları simüle ederek en kısa yolları bulur ve sonuçları kullanıcıya sunar.

.Kullanım
Programı çalıştırmak için terminalde şu komutu kullanabilirsiniz:
go run . istenilen.txt
Burada, istenilen.txt dosyası grafın ve diğer gerekli bilgilerin bulunduğu dosyadır. Program bu dosyayı okur, grafı oluşturur ve en kısa yolları bulur.

.Girdi Dosyası Formatı
Girdi dosyasının belirli bir formatı vardır. Örnek bir girdi dosyası şu şekildedir:
3
A 0 0
B 1 2
C 3 4
##start
A
##end
C
A-B
B-C
Bu dosya, üç düğüm içerir (A, B ve C) ve bunların koordinatlarını tanımlar. Ayrıca, başlangıç düğümü A ve bitiş düğümü C olarak belirlenmiştir. Bağlantılar da A-B ve B-C şeklinde tanımlanmıştır.

.Örnek Çıktı
Program çalıştırıldıktan sonra, en kısa yollar ve diğer ilgili bilgiler terminalde görüntülenir. Örnek bir çıktı şu şekildedir:
Karınca sayısı: 3
Koordinatlar: map[A:[0 0] B:[1 2] C:[3 4]]
Başlangıç Koordinatı: A
Bitiş Koordinatı: C
Bağlantılar: [A-B B-C]
Graf oluşturuldu: Nodes: [A [0 0] B [1 2] C [3 4]] Edges: [[A B] [B C] [C B] [B A]]
Tüm kısa yollar:
Path 0: A B C
Seçilen Yollar (Her Başlangıç İndeksi İçin Minimum Sayı):
Path 0: [A B C]
Benzersiz Yollar (Her Bitiş Düğümü İçin, Boş Yollar Dahil):
Path 0: [A B C]
Bitiş Düğümü Eklenmiş Benzersiz Yollar:
Path 0: [A B C]
Düğümler Olarak Bitiş Düğümü Eklenmiş Benzersiz Yollar:
Path 0: A B C
Bu çıktı, programın çalıştırılması ve sonuçların görüntülenmesi sırasında elde edilen bilgileri içerir.