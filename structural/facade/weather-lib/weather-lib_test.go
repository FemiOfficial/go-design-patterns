package weatherlib

import (
	"fmt"
	"godesignpatterns/config"
	"godesignpatterns/structural/facade/weather-lib/utils"
	"testing"

	"github.com/joho/godotenv"
)
func TestWeatherApi(t *testing.T) {
	apiKey := "1234567890"

	mockData := utils.GetMockWehaterData()
	weatherApi := CurrentWeatherData{
		APIKey: apiKey,
	}

	weather, err := weatherApi.ResponseParser(mockData)
	if err != nil {
		t.Fatal(err)
	}

	if weather.Name != "Province of Turin" {
		t.Errorf("Weather name should be Province of Turin")
	}
}

func TestWeatherApiByCityAndCountryCode(t *testing.T) {

	// using .env file to set the API key for the weather API
	// might consider a better approach in the future
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Fatal(err)
	}	

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	apiKey := cfg.OpenWeatherAPIKey
	if apiKey == "" {
		t.Fatal("OPENWEATHER_API_KEY is not set")
	}
	weatherApi := CurrentWeatherData{
		APIKey: apiKey,
	}
	weather, err := weatherApi.GetByCityAndCountryCode("ibadan", "NG")
	if err != nil {
		t.Fatal(err)
	}

	if (err != nil) {
		t.Errorf("Error should be nil, weather api should be working")
	}

	fmt.Println("current weather temperature for ibadan, NG: ", weather.Main.Temp)
}