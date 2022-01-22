package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openw/shared"
)

var url = "https://api.openweathermap.org/data/2.5/onecall"

type OpenWeatherResponse struct {
	Current struct {
		FeelsLike float64 `json:"feels_like"`
		Weather   []struct {
			Main string `json:"main"`
		} `json:"weather"`
	} `json:"current"`
	Alerts []struct {
		Event string `json:"event"`
	} `json:"alerts"`
}

type OpenWeatherDal struct{}

func (openWeatherDal OpenWeatherDal) GetWeather(lat string, lon string, appid string) (shared.WeatherResposne, error) {
	params := fmt.Sprintf("lat=%s&lon=%s&appid=%s&units=imperial", lat, lon, appid)
	path := fmt.Sprintf("%s?%s", url, params)

	resp, err := http.Get(path)

	if err != nil {
		return shared.WeatherResposne{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return shared.WeatherResposne{}, err
	}

	var owResp = OpenWeatherResponse{}
	err = json.Unmarshal(body, &owResp)

	if err != nil {
		return shared.WeatherResposne{}, err
	}

	return shared.WeatherResposne{
		Condition: owResp.Current.Weather[0].Main,
		Temp:      GetTempVerbiage(owResp.Current.FeelsLike),
		HasAlert:  len(owResp.Alerts) > 0,
		Alerts:    GetAlertStringsFromResp(owResp),
	}, err
}

func GetAlertStringsFromResp(owResp OpenWeatherResponse) []string {
	if len(owResp.Alerts) < 1 {
		return nil
	}

	var alerts []string

	for _, alert := range owResp.Alerts {
		alerts = append(alerts, alert.Event)
	}

	return alerts
}

func GetTempVerbiage(feelsLike float64) string {
	switch {
	case feelsLike > -150 && feelsLike < 10:
		return shared.VeryColdString
	case feelsLike < 40:
		return shared.ColdString
	case feelsLike < 60:
		return shared.TolerableString
	case feelsLike < 80:
		return shared.NiceString
	case feelsLike >= 80 && feelsLike < 150:
		return shared.HotString
	default:
		return "Unable to get temp"
	}
}
