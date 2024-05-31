package apps

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)


func PttCrawler(board string) []*discordgo.MessageEmbed {
	var articles []*discordgo.MessageEmbed
	baseUrl := "https://www.ptt.cc/"
	url := fmt.Sprintf("https://www.ptt.cc/bbs/%s/index.html", board)
	fmt.Printf("Now parsing url: %s\n", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: "over18", Value: "1"})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status code is: %d\n", resp.StatusCode)

	if resp.StatusCode != 200 {
		log.Fatalf("Status code: %d, %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var count int

	doc.Find("#main-container > div.r-list-container.action-bar-margin.bbs-screen").Find(".r-ent, .r-list-sep").EachWithBreak(
		func(i int, s *goquery.Selection) bool {
			if s.HasClass("r-list-sep") || count == 10 {
				return false
			}
			titles := s.Find(".title")
			
			deletedTitle := titles.Text()
			if strings.Contains(deletedTitle, "本文已被刪除") {
				return true
			}
			titleA := titles.Find("a")
			title := strings.TrimSpace(titleA.Text())
			
			articleUrl, _ := titleA.Attr("href")
			article := &discordgo.MessageEmbed{
					Title: fmt.Sprintf("%d. %s", count+1, title),
					URL: fmt.Sprintf("%s%s", baseUrl, articleUrl),
			}
			// fmt.Println(title)
			articles = append(articles, article)
			count++
			return true
		})
	
	return articles
}