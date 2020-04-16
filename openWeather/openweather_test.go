package openWeather

import (
	"context"
	"currentWeatherBot/models"
	"fmt"
	"log"
	"testing"
)

func TestGetCurrentWeatherByCityName(t *testing.T) {

	extraURL := "/data/2.5/weather"

	ctx := context.Background()

	Init("key", "url")

	var currentWeather models.CurrentWeather

	param := map[string]string{
		"q": "Selangor",
	}

	err := Get(ctx, extraURL, param, &currentWeather)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentWeather)

}

func TestGetCurrentWeatherByLatitudeAndLongitude(t *testing.T) {

	extraURL := "/data/2.5/weather"

	ctx := context.Background()

	Init("key", "url")

	var currentWeather models.CurrentWeather

	param := map[string]string{
		"lon": "101.6428922",
		"lat": "2.9266403",
	}

	err := Get(ctx, extraURL, param, &currentWeather)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentWeather)

}
