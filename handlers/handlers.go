package handlers

import (
	"context"
	"currentWeatherBot/helpers"
	"currentWeatherBot/openWeather"
	"fmt"
	"github.com/codingsince1985/geo-golang/locationiq"
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

var (
	START_RESOURCE = `
Current Weather Updates ☀️ ☁️ 🌩️ 🌧 
#ItsWeatherTime

Hello %s ❗👋
This bot provides weather information instantly 🙌
👉 /city 
	➡️ Get city weather
👉 /my_location  
	➡️ Get updates for your location`

	LOCATION_RESOURCE = `
%s's Weather Now ! ☀️
👉 Weather Condition : 
	   		%s (%s)
👉 Temperature										 : %.2f °C
👉 Min.Temperature 		 : %.2f °C
👉 Max.Temperature 	 : %.2f °C
👉 Humidity            			   : %d  K
👉 Country  																  : %s
`
)

func Init(bot *telebot.Bot, locationiqToken string) *telebot.Bot {

	ctx := context.Background()

	bot.Handle("/start", func(m *telebot.Message) {
		bot.Send(m.Sender, fmt.Sprintf(START_RESOURCE, m.Sender.FirstName))
	})

	bot.Handle("/city", func(m *telebot.Message) {
		bot.Send(m.Sender, fmt.Sprintf(`Please enter name of city 🏙️`))
	})

	bot.Handle(telebot.OnText, func(q *telebot.Message) {

		city := q.Text
		weather, err := openWeather.GetCurrentWeatherByCityName(ctx, city)
		if err != nil {
			log.Println(err)
		}

		if weather.Sys.Country == "" {
			bot.Send(q.Sender, fmt.Sprintf(`Incorrect input ❌ , Please enter again ❗`))

		} else {
			if len(weather.Weather) != 0 {
				bot.Send(q.Sender, fmt.Sprintf(LOCATION_RESOURCE, city, weather.Weather[0].Main,
					weather.Weather[0].Description, helpers.KelvinToCelcius(weather.Main.Temp),
					helpers.KelvinToCelcius(weather.Main.TempMin), helpers.KelvinToCelcius(weather.Main.TempMax),
					weather.Main.Humidity, weather.Sys.Country), &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
			} else {
				bot.Send(q.Sender, fmt.Sprintf(`Try Again ❗`))
			}
		}
	})

	bot.Handle("/my_location", func(m *telebot.Message) {
		bot.Send(m.Sender, fmt.Sprintf(`Please send your location 📍`))
	})

	bot.Handle(telebot.OnLocation, func(q *telebot.Message) {
		location, err := locationiq.Geocoder(locationiqToken, 18).ReverseGeocode(float64(q.Location.Lat),
			float64(q.Location.Lng))
		if err != nil {
			log.Println(err)
		}

		if location == nil {
			bot.Send(q.Sender, `Sorry, your location not found`, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		} else {
			weather, err := openWeather.GetCurrentWeatherByCityName(ctx, location.City)
			if err != nil {
				log.Println(err)
			}
			bot.Send(q.Sender, fmt.Sprintf(LOCATION_RESOURCE, location.City, weather.Weather[0].Main,
				weather.Weather[0].Description, helpers.KelvinToCelcius(weather.Main.Temp),
				helpers.KelvinToCelcius(weather.Main.TempMin), helpers.KelvinToCelcius(weather.Main.TempMax),
				weather.Main.Humidity, weather.Sys.Country), &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})

		}
	})
	return bot
}
