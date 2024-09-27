package mmongodb

type WeatherData struct {
	CityName string `json:"name"`
	Current  struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}
