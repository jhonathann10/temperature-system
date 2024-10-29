package weatherapi

type WeatherAPIInterface interface {
	GetWeatherByCity(city string) (*Weather, error)
}
