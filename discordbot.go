package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
)

const (
	token = "ODQ4ODAxMzYzMzY5Nzg3NDQ0.YLR53g.mdfjzO1DyFow8d5D_uDdP8rcomQ"
)

func startBot() (*discordgo.Session, error) {
	log.Println("!-STARTING BOT-!")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Invalid bot parameters: %v", err)
	}

	err = s.Open()
	if err != nil {
		return nil, fmt.Errorf("Error connecting: %v", err)
	}
	log.Println("!-BOT STARTED-!")
	return s, nil
}

func registerCommands(r *dgc.Router) {
	r.RegisterCmd(&dgc.Command{
		Name:        "hello",
		Description: "Says hello",
		Usage:       "hello",
		Example:     "hello",
		IgnoreCase:  true,
		Handler: func(ctx *dgc.Ctx) {
			err := ctx.RespondText("Hello there!")
			if err != nil {
				log.Fatalf("hello: %v", err)
			}
		},
	})
	r.RegisterCmd(&dgc.Command{
		Name:        "solarian",
		Description: "get solarian by mint number",
		Usage:       "solarian",
		Example:     "solarian 1",
		IgnoreCase:  true,
		Handler: func(ctx *dgc.Ctx) {
			solarian := ctx.Arguments.Get(0).Raw()
			if len(solarian) == 1 {
				err := ctx.RespondText("You requested solarian " + solarian)
				if err != nil {
					log.Fatalf("solarian: %v", err)
				}
			} else if len(solarian) > 1 {
				embed := &discordgo.MessageEmbed{
					Type: discordgo.EmbedTypeGifv,
					Image: &discordgo.MessageEmbedImage{
						URL: "https://solarians.click/render/" + solarian + ".gif",
					},
				}
				err := ctx.RespondTextEmbed("Here is solarian mint: "+solarian, embed)
				if err != nil {
					log.Fatalf("solarian: %v", err)
				}
			}
		},
	})
}

func main() {
	session, err := startBot()
	if err != nil {
		panic(err)
	}
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	registerCommands(router)

	router.Initialize(session)

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
	}()
}
