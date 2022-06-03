package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/emrullahcirit/weather-cli/model"
	"github.com/joho/godotenv"
)

func getRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	return body, err
}

func UnixToDate(unixTime int) string {
	return time.Unix(int64(unixTime), 0).Format("02-01-2006")
}

func GetCurrentWeather(cityName string) (model.CurrentWeather, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")

	openWeatherURL := "https://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&lang=tr" + "&appId=" + apiKey + "&units=metric"

	resp, err := getRequest(openWeatherURL)

	if err != nil {
		log.Fatalln(err)
	}

	var currentWeather model.CurrentWeather

	err = json.Unmarshal(resp, &currentWeather)

	return currentWeather, err
}

func PrintWeatherInfos(city string) {
	w, err := GetCurrentWeather(city)

	if err != nil {
		log.Fatalln(err)
	}

	humidity := strconv.Itoa(w.Forecast.Humidity)
	date := UnixToDate(w.Date)
	infos := fmt.Sprintf("Şehir: %s, %s\n\nTarih: %s\n\nSıcaklık: %.2f °C\n\nHava Durumu: %s\nHissedilen: %.2f °C\nMaximum Sıcaklık: %.2f °C\nMinimum Sıcaklık: %.2f °C\nNem: %s\nRüzgarın Hızı: %.2f", w.CityName, w.Info.Country, date, w.Forecast.Temp, strings.Title(w.Weather[0].Description), w.Forecast.FeelsLike, w.Forecast.Max, w.Forecast.Min, humidity, w.Wind.Speed)

	Box := box.New(box.Config{Px: 6, Py: 3, Type: "Bold", TitlePos: "Top", Color: "White", ContentAlign: "Left"})
	Box.Print("Hava Durumu Bilgisi", infos)

}
