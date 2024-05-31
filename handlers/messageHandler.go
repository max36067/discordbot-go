package handlers

import (
	"golang-discord-bot/v2/apps"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	boardButtonsComponent := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label: "Gossiping",
					Style: discordgo.PrimaryButton,
					Disabled: false,
					CustomID: "Gossiping",
				},
				discordgo.Button{
					Label: "C_Chat",
					Style: discordgo.PrimaryButton,
					Disabled: false,
					CustomID: "C_Chat",
				},
				discordgo.Button{
					Label: "Soft_Job",
					Style: discordgo.PrimaryButton,
					Disabled: false,
					CustomID: "Soft_Job",
				},
			},
		}}

	switch {
	case m.Content == "!ptt":
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "請選擇看板",
			Components: boardButtonsComponent,
		})

	case strings.Contains(m.Content, "!pic"):
		magicString := strings.Split(m.Content, "-")
		image := apps.PictureGenerate(magicString[1])
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			File: &discordgo.File{
				Name: "image.jpeg",
				ContentType: "image/jpeg",
				Reader: image,
			},
			Content: "已為您產生圖片",
		})

	} 
}