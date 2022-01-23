package shared

var VeryColdString = "Very Cold"
var ColdString = "Cold"
var TolerableString = "Tolerable"
var NiceString = "Nice"
var HotString = "Hot"

type WeatherResposne struct {
	Condition string   `json:"condition"`
	Temp      string   `json:"temp"`
	HasAlert  bool     `json:"hasAlert"`
	Alerts    []string `json:"alerts"`
}
