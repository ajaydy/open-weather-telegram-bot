package main

import (
	config2 "currentWeatherBot/config"
	"currentWeatherBot/handlers"
	"currentWeatherBot/openWeather"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func main() {

	config, err := config2.Init(".config.toml")
	if err != nil {
		log.Fatal(err)
	}
	openWeather.Init(config.ApiKey, config.OpenweatherURL)

	webhook := &tb.Webhook{
		Listen: ":" + config.Port,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: config.PublicUrl,
		},
	}

	pref := tb.Settings{
		Token:  config.Token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	handlers.Init(b, config.LocationiqToken)
	b.Start()
}
