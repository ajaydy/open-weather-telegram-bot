package openWeather

import (
	"context"
	"currentWeather/models"
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {

	ctx := context.Background()

	extraURL := "/data/2.5/weather"

	Init("key", "url")

	param := map[string]string{
		"q": "Selangor",
	}

	var currentWeatherResponse models.CurrentWeather

	err := Get(ctx, extraURL, param, &currentWeatherResponse)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentWeatherResponse)

}
