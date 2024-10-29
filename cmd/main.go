package main

import (
	"fmt"

	"github.com/jhonathann10/temperature-system/configs"
	"github.com/jhonathann10/temperature-system/internal/infra"
	"github.com/jhonathann10/temperature-system/internal/infra/client/viacep"
	"github.com/jhonathann10/temperature-system/internal/infra/client/weatherapi"
	"github.com/jhonathann10/temperature-system/internal/infra/webserver"
)

const (
	baseURLViaCEP  = "https://viacep.com.br/ws/"
	baseURLWeather = "http://api.weatherapi.com/v1"
	serverPort     = ":8080"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	server := webserver.NewWebServer(serverPort)
	cepHandler := viacep.NewViaCEPClient(baseURLViaCEP)
	weatherHandler := weatherapi.NewWeatherAPIClient(baseURLWeather, config.KeyAPIWeatherApi)
	handler := infra.NewHandler(cepHandler, weatherHandler)
	server.AddHandler("/temperature", handler.GetTemperature)
	fmt.Println("Starting web server on port", serverPort)

	server.Start()
}
