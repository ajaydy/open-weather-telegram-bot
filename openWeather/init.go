package openWeather

var (
	ApiKey string
	URL    string
)

func Init(key string, url string) {
	ApiKey = key
	URL = url
}
