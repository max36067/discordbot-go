package main

import (
	"flag"
	"fmt"
	"golang-discord-bot/v2/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token       string
	GuildID		string
)

func init() {
	flag.StringVar(&GuildID, "guild", "", "GuildID")
	flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}

func main() {
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }
	s.AddHandler(handlers.MessageCreate)
	s.AddHandler(handlers.CrawlerHandler)

	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	defer s.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("Removing commands...")


	// Cleanly close down the Discord session.
	s.Close()
}

