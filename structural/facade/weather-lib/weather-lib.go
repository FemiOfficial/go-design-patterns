package weatherlib

import (
	"encoding/json"
	"fmt"
	"godesignpatterns/structural/facade/weather-lib/model"
	"io"
	"net/http"
	"os"
)

type CurrentWeatherData struct {
	APIKey string
}
	
type CurrentWeatherRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (model.Weather, error)
	GetByCoordinates(latitude, longitude float64) (model.Weather, error)
}

func (c *CurrentWeatherData) ResponseParser(data io.Reader) (*model.Weather, error) {
	weather := new(model.Weather)
	err := json.NewDecoder(data).Decode(weather)
	return weather, err
}

func (c* CurrentWeatherData) DoRequest(url string) (*model.Weather, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stdout, "Request Statuscode %d \n", response.StatusCode)
	fmt.Fprintf(os.Stdout, "Request ContentLength %d \n", response.ContentLength)
	fmt.Fprintf(os.Stdout, "Request Host %s \n", response.Request.Host)
	fmt.Fprintf(os.Stdout, "Request Method %s \n", response.Request.Method)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", response.StatusCode)
	}

	weather, err := c.ResponseParser(response.Body)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (*model.Weather, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", city, countryCode, c.APIKey)
	return c.DoRequest(url)
}

func (c *CurrentWeatherData) GetByCoordinates(latitude, longitude float64) (*model.Weather, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", latitude, longitude, c.APIKey)
	return c.DoRequest(url)
}