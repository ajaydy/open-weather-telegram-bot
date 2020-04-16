package openWeather

import (
	"context"
	"currentWeatherBot/models"
)

func GetCurrentWeatherByCityName(ctx context.Context, city string) (models.CurrentWeather, error) {

	extraURL := "/data/2.5/weather"

	var currentWeather models.CurrentWeather

	param := map[string]string{
		"q": city,
	}

	err := Get(ctx, extraURL, param, &currentWeather)
	if err != nil {
		return models.CurrentWeather{}, nil
	}

	return currentWeather, nil
}

func GetCurrentWeatherByLatitudeAndLongitude(ctx context.Context, longitude string, latitude string) (
	models.CurrentWeather, error) {

	extraURL := "/data/2.5/weather"

	var currentWeather models.CurrentWeather

	param := map[string]string{
		"lon": longitude,
		"lat": latitude,
	}

	err := Get(ctx, extraURL, param, &currentWeather)
	if err != nil {
		return models.CurrentWeather{}, nil
	}

	return currentWeather, nil
}
