package main

import (
	"fmt"
	"github.com/Feyd-Rauth/Zlyden/pkg/zlyden"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// loading Discord bot token
	token, err := loadToken("token.dat")
	if err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(zlyden.MessageCreate)

	if err := discord.Open(); err != nil {
		log.Fatal(err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(discord)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func loadToken(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
