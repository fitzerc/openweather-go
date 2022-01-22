package shared

type OpenwDal interface {
	GetWeather(lat string, lon string, appid string) (WeatherResposne, error)
}
