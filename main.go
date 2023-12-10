package main

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	color "github.com/gookit/color"
	"time"
)

func main() {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(formatForDisplay(&message))
	})

	client.Join("myshkovsky")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func formatForDisplay(message *twitch.PrivateMessage) string {
	return fmt.Sprintf(
		"%s %s %s: %s",
		formatTimestamp(),
		formatBadges(&message.User.Badges),
		formatName(&message.User.Name, &message.User.Color),
		formatMessage(&message.Message),
	)
}

func formatTimestamp() string {
	h, m, s := time.Now().Clock()
	return fmt.Sprintf("[%02d:%02d:%02d]", h, m, s)
}

func formatName(name *string, colorHex *string) string {
	return color.HEX(*colorHex).Sprint(*name)
}

func formatMessage(s *string) string {
	return *s
}

func formatBadges(badges *map[string]int) string {
	s := ""
	for k := range *badges {
		switch k {
		case "admin", "staff":
			s += "👮"
		case "broadcaster":
			s += "🌌"
		case "moderator":
			s += "🛡"
		case "subscriber":
			s += "🌟"
		}
	}
	return s
}
