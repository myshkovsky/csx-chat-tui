package main

import (
	"embed"
	"errors"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/dolmen-go/kittyimg"
	"github.com/gempir/go-twitch-irc/v4"
	color "github.com/gookit/color"
)

var wg *sync.WaitGroup
var cachePath string
//go:embed cache/emotes/*
var files embed.FS

func init() {
    wg = new(sync.WaitGroup)
    f, err := os.Lstat("./")
    cachePath = "cache/emotes"
    err = os.MkdirAll(cachePath, f.Mode().Perm())
    catchError(err)
}

func main() {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(formatForDisplay(&message))
	})

	client.Join("myshkovsky")

	err := client.Connect()
	catchError(err)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func formatForDisplay(message *twitch.PrivateMessage) string {
	return fmt.Sprintf(
		"%s %s %s: %s",
		formatTimestamp(),
		formatBadges(&message.User.Badges),
		formatName(&message.User.Name, &message.User.Color),
		formatMessage(&message.Message, &message.Emotes),
	)
}

func formatTimestamp() string {
	h, m, s := time.Now().Clock()
	return fmt.Sprintf("[%02d:%02d:%02d]", h, m, s)
}

func formatName(name *string, colorHex *string) string {
	return color.HEX(*colorHex).Sprint(*name)
}

func formatMessage(s *string, emotes *[]*twitch.Emote) string {
	// TODO: Format Twitch emotes and display them
	// See: printEmote
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

func formatEmotes(msg *string, emotes *[]*twitch.Emote) string {
	formatted := *msg
	for _, emote := range *emotes {
		for _, v := range emote.Positions {
			formatted = formatted[:v.Start] + "<@>" + emote.Name + "<@>" + formatted[v.End+1:]
		}
	}
	return formatted
}

// WARNING: Not implemented: Most Windows terminals are unable to display the image
func printEmote(id string, name string) {
	if !fileExists(fmt.Sprintf("%s/%s.png", cachePath, name)) {
		wg.Add(1)
		go downloadImg(id, name)
	}
	wg.Wait()
	f, err := files.Open(fmt.Sprintf("%s/%s.png", cachePath, name))
	defer f.Close()
	catchError(err)
	img, err := png.Decode(f)
	catchError(err)
	kittyimg.Fprint(os.Stdout, img)
}

func downloadImg(id string, name string) {
	defer wg.Done()
	url := fmt.Sprintf(
		"https://static-cdn.jtvnw.net/emoticons/v2/%s/default/dark/1.0",
		id,
	)
	path := fmt.Sprintf("%s/%s.png", cachePath, name)
	err := exec.Command("curl", url, "-o", path).Run()
	catchError(err)
}

func catchError(err error) {
	panic(err)
}
