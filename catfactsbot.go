package main

import (
	"fmt"
	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"os"
	curl "github.com/andelf/go-curl"
	"encoding/json"
)

func main() {
	ListenForRequest()
}

func ListenForRequest() {
	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
	bot.Hear("(?i)catfact(.*)").MessageHandler(CatFactHandler)
	bot.Run()
}

func CatFactHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://catfacts-api.appspot.com/api/facts?number=1")
	fooTest := func (buf []byte, userdata interface{}) bool {
		var dat map[string]interface{}
		if err := json.Unmarshal(buf, &dat); err != nil {
			panic(err)
	    }
	    strs := dat["facts"].([]interface{})
    	fact := strs[0].(string)
		bot.Reply(evt, fmt.Sprintln(fact), slackbot.WithTyping)
		return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

    if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
    }
}
