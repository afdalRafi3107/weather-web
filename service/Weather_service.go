package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

type WeatherAPIResponse struct {
	NAME string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"tempt_max"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`

	Weather []struct {
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"wether"`

	Dt       int64 `json:"dt"`
	TimeZOne int64 `json:"time_zome"`
}

func FetchCurrentWeather(city, units string) (*WeatherAPIResponse, error) {
	base := viper.GetString("OWM_BASE_URL")

	if base == ""{
		base = "https://api.openweathermap.org/data/2.5"
	}

	key := viper.GetString("OWM_API_KEY")

	if key == ""{
		return  nil, errors.New("API not found")
	}

	if units == "" {
		units = "metric"
	}

	u := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=%s",
	base,
	url.QueryEscape(city),
	url.QueryEscape(key),
	url.QueryEscape(units),
)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(u)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300{
		return  nil, fmt.Errorf("weather API Error : %s", resp.Status)
	}

	var out WeatherAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil{
		return nil, err
	}
	
	return  &out, nil
}
