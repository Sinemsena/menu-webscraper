# Menü ile Çalışan Web Scraper

Golang kullanarak web sitelerinden başlık, açıklama ve tarih bilgilerini çeken CLI web

scraper aracıdır.



## Kullanım



1. Program Başlatma:

Herhangi bir filtre uygulamadan çalıştırdığınızda seçim menüsü görüntülenerek bir seçim yapmanız istenir:



```bash



go run main.go

```

Program başladığında karşınıza gelen seçim menüsünden -1, -2 veya -3 girerek taramayı başlatabilir, -4 ile çıkış yapabilirsiniz.



2. Filtreleme Seçenekleri:

Filtreler programı başlatırken girilir ve o oturum boyunca geçerli olur:



* Tarih Filtreleme: Tarih bilgisini gizlemek için -date kullanılır ve daha sonra menüden seçim yapılır.



```bash



go run main.go -date

```

 * Açıklama Filtreleme: Açıklama bilgisini gizlemek için -description kullanılır ve daha sonra menüden seçim yapılır.



```bash



go run main.go -description

```

 * Tam Filtreleme: Sadece haber başlıklarını görmek için her iki parametreyi de ekleyerek filtreleme yapılır ve daha sonra menüden seçim yapılır.



```bash



go run main.go -date -description

```

## Kayıt Sistemi

Tarama Sonuçları, projenin kök dizininde otomatik olarak oluşturulan /sonuc klasörüne, o anki saat damgasıyla .txt formatında kaydediliyor.
