package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func umlToImage(uml string) (*os.File, error) {
	uid := uuid.New().String()
	umlFileName := "uml/" + uid + ".pu"
	umlFile, err := os.Create(umlFileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = umlFile.Close()
	}()

	if _, err := umlFile.Write([]byte(uml)); err != nil {
		return nil, err
	}

	cmd := exec.Command("java", "-jar", "plantuml.jar", "-o", "out", umlFileName)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return os.Open("uml/out/" + uid + ".png")
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "```uml") && strings.HasSuffix(m.Content, "```") {
		image, err := umlToImage(m.Content[6 : len(m.Content)-3])
		if err != nil {
			log.Println(err)
			_, err := s.ChannelMessageSend(m.ChannelID, "Failed to convert uml")
			if err != nil {
				log.Println(err)
			}
			return
		}
		_, err = s.ChannelFileSend(m.ChannelID, "uml.png", image)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid arguments length")
	}

	session, err := discordgo.New()
	if err != nil {
		log.Fatal(err)
	}

	session.Token = "Bot " + os.Args[1]
	session.AddHandler(onMessage)
	if err := session.Open(); err != nil {
		log.Fatal(err)
	}

	log.Println("uml2image discord bot started")
	<-make(chan bool)
}
