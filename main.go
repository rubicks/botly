package main

import (
	"fmt"
	slackbot "github.com/rubicks/botly/Godeps/_workspace/src/github.com/BeepBoopHQ/go-slackbot"
	"github.com/rubicks/botly/Godeps/_workspace/src/github.com/briandowns/openweathermap"
	"github.com/rubicks/botly/Godeps/_workspace/src/github.com/nlopes/slack"
	"github.com/rubicks/botly/Godeps/_workspace/src/golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("main")
	ListenForWeather()
}

func GetWeather(place string) (*openweathermap.CurrentWeatherData, error) {
	log.Println("GetWeather")
	w, err := openweathermap.NewCurrent("F", "en")
	if err != nil {
		return nil, fmt.Errorf("Could not get weather: %v", err)
	}

	err = w.CurrentByName(place)
	if err != nil {
		return nil, fmt.Errorf("Weather fetch fail: %v", err)
	}
	return w, nil
}

func ListenForWeather() {
	log.Println("ListenForWeather")
	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
	bot.Hear("(?i)weather (.*)").MessageHandler(WeatherHandler)
	bot.Run()
}

func WeatherHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	log.Println("WeatherHandler")
	parts := strings.Split(evt.Msg.Text, " ")
	if len(parts) != 2 {
		return
	}
	weather, err := GetWeather(parts[1])
	if err != nil {
		fmt.Println("Could not get weather:", err)
		return
	}
	description := ""
	if len(weather.Weather) > 0 {
		description = weather.Weather[0].Description
	}
	bot.Reply(evt,
		fmt.Sprintf("The current temperature for %s is %.0f degrees farenheight (%s)",
			weather.Name,
			weather.Main.Temp,
			description),
		slackbot.WithTyping)
}
