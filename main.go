package main

import (
	"bufio"
	"flag"
	"fmt"
	"time"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/mbndr/figlet4go"
)

func main() {
	ascii, _ := figlet4go.NewAsciiRender().Render("Yavuzlar-2")
	color.Cyan(ascii)
	tarihfiltreli := flag.Bool("date", false, "filters the date part")
	aciklamafiltreli := flag.Bool("description", false, "filters the description part")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(color.YellowString("\n   *SEÇİM MENÜSÜ*   "))
		fmt.Println(" -1 -> The Hacker News\n -2 -> Security Week\n -3 -> Hack Read\n -4 -> Çıkış")
		fmt.Print(color.CyanString("\nSeçiminiz: "))

		scanner.Scan()
		secim := strings.TrimSpace(scanner.Text())

		url, siteAdi := "", ""
		switch secim {
		case "-1":
			url, siteAdi = "https://thehackernews.com/", "HackerNews"
		case "-2":
			url, siteAdi = "https://www.securityweek.com/", "SecurityWeek"
		case "-3":
			url, siteAdi = "https://www.hackread.com/", "HackRead"
		case "-4":
			color.Green("\nÇıkış yapılıyor")
			return
		default:
			color.Red("Geçersiz seçim lütfen -1, -2, -3 veya -4 gir")
			continue
		}
		scrapeAndShow(url, siteAdi, *tarihfiltreli, *aciklamafiltreli)
	}}

func scrapeAndShow(url string, siteAdi string, tarihFiltreleme bool, aciklamaFiltereleme bool) {
	klasörAdi :="sonuc"
	//klasör yoks olştur
	if _, err := os.Stat(klasörAdi); os.IsNotExist(err) {
		err := os.Mkdir(klasörAdi, 0755)
			if err != nil {
				color.Red("Klasör oluşturulamadı", err)
				return
			}
		}

	// Her çalışma için ayrı dosya
	dosyaAdi := fmt.Sprintf("%s/%s_%s", klasörAdi, siteAdi, time.Now().Format("150405")) 
	if tarihFiltreleme { dosyaAdi += "_tarihfiltreli" }
	if aciklamaFiltereleme { dosyaAdi += "_aciklamafiltreli" }
	dosyaAdi += ".txt"
	file, _ := os.Create(dosyaAdi)
	defer file.Close()
	writer := bufio.NewWriter(file)

	c := colly.NewCollector()
	sayac := 1

	c.OnHTML("article, .body-post, .cat-post-item", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildText("h2, .home-title, .cs-entry__title"))
		if title == "" { return }
		color.Magenta("%d. Haber:", sayac)
		fmt.Println(title)
		writer.WriteString(fmt.Sprintf("%d. Haber: %s\n", sayac, title))

		if !aciklamaFiltereleme {
			aciklama := strings.TrimSpace(e.ChildText("p, .home-desc, .cs-entry__excerpt"))
			if aciklama != "" {
				color.Blue("Açıklama:")
				fmt.Println(aciklama)
				writer.WriteString("Açıklama: " + aciklama + "\n")
			}
		}

		if !tarihFiltreleme {
			tarih := strings.TrimSpace(e.ChildText("time, .h-datetime, .cs-meta-date"))
			if tarih != "" {
				color.Red("Tarih:")
				fmt.Println(tarih)
				writer.WriteString("Tarih: " + tarih + "\n")
			}
		}
		fmt.Println("") // Haberler arası 1 satır boşluk
		writer.WriteString("\n")
		sayac++
	})
	c.Visit(url)
	writer.Flush()
	color.Green("\n %d haber '%s' dosyasına kaydedildi.", sayac-1, dosyaAdi)
}