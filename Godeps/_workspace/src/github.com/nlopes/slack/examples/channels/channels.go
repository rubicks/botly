package main

import (
	"fmt"

	"github.com/rubicks/botly/Godeps/_workspace/src/github.com/nlopes/slack"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.ID)
	}
}
