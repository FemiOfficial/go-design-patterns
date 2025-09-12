package model
// Weather represents the complete weather data structure
type Weather struct {
	Coord      Coordinates `json:"coord"`
	Weather    []WeatherInfo `json:"weather"`
	Base       string     `json:"base"`
	Main       Main       `json:"main"`
	Visibility int        `json:"visibility"`
	Wind       Wind       `json:"wind"`
	Rain       *Rain      `json:"rain,omitempty"`
	Clouds     Clouds     `json:"clouds"`
	Dt         int64      `json:"dt"`
	Sys        Sys        `json:"sys"`
	Timezone   int        `json:"timezone"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Cod        int        `json:"cod"`
}

// Coordinates represents geographical coordinates
type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// WeatherInfo represents weather condition details
type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main represents main weather parameters
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

// Wind represents wind information
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

// Rain represents rain information
type Rain struct {
	OneHour float64 `json:"1h"`
}

// Clouds represents cloud information
type Clouds struct {
	All int `json:"all"`
}

// Sys represents system information
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}