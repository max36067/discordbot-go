package handlers

import (
	"encoding/json"
	"fmt"
	"golang-discord-bot/v2/apps"

	"github.com/bwmarrin/discordgo"
)




func CrawlerHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	boardId := i.MessageComponentData().CustomID
	fmt.Printf("Now parsing '%s' board\n", boardId)
	articles := apps.PttCrawler(boardId)
	jsonArticles, _ := json.Marshal(articles)
	fmt.Println(string(jsonArticles))

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("以下是 %s 看板的文章：", boardId),
			Embeds: articles,
		},
	}

	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
		
	
}