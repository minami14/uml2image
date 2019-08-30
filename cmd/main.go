package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func trimSpaceAndNewLine(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, "\r")
	return s
}

func formatUml(uml string) string {
	uml = trimSpaceAndNewLine(uml)

	if !strings.HasPrefix(uml, "@startuml") {
		uml = "@startuml\n" + uml
	}

	uml = trimSpaceAndNewLine(uml)

	if !strings.HasSuffix(uml, "@enduml") {
		uml = uml + "\n@enduml"
	}

	return uml
}

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

	uml = formatUml(uml)

	if _, err := umlFile.Write([]byte(uml)); err != nil {
		return nil, err
	}

	cmd := exec.Command("java", "-jar", "plantuml.jar", "-o", "out", umlFileName)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return os.Open("uml/out/" + uid + ".png")
}

const (
	prefix = "```uml"
	suffix = "```"
)

func isUml(s string) bool {
	return strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix)
}

func extractUml(s string) string {
	return s[len(prefix) : len(s)-len(suffix)]
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	text := trimSpaceAndNewLine(m.Content)
	if isUml(text) {
		uml := extractUml(text)
		image, err := umlToImage(uml)
		defer func() {
			_ = image.Close()
		}()

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
