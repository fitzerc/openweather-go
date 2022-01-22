package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openw/openweather"
	"openw/shared"
	"os"

	"github.com/gin-gonic/gin"
)

var dal shared.OpenwDal
var configuration Configuration

func main() {
	//TODO: abstract config read to make source flexible
	var err error
	configuration, err = GetConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	if configuration.AppId == "Add api key here" {
		log.Fatal("Please add your api key to config/config.json")
	}

	dal = openweather.OpenWeatherDal{}
	router := gin.Default()

	router.GET("/weather", getWeather)

	router.Run("localhost:" + configuration.Port)
}

func getWeather(c *gin.Context) {
	//TODO: add error handling for bad request
	lat := c.Query("lat")
	lon := c.Query("lon")

	currentWeather, err := dal.GetWeather(lat, lon, configuration.AppId)

	//TODO: Add custom errors and return appropriate status
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("{\"error\": %s}", err.Error()))
	} else {
		c.IndentedJSON(http.StatusOK, currentWeather)
	}
}

func GetConfig() (Configuration, error) {
	//TODO: abstract config read to make source flexible
	file, err := os.Open("./config/config.json")

	if err != nil {
		return Configuration{}, err
	}

	decoder := json.NewDecoder(file)

	var config Configuration
	err = decoder.Decode(&config)

	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}

type Configuration struct {
	AppId string
	Port  string
	//TODO: Allow use of other supported units
	//Units string
}
