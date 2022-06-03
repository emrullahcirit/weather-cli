package model

type Wind struct {
	Speed float64 `json:"speed"`
}

type Forecast struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Min       float64 `json:"temp_min"`
	Max       float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type CurrentWeather struct {
	Weather  []Weather `json:"weather"`
	Forecast Forecast  `json:"main"`
	Wind     Wind      `json:"wind"`
	CityName string    `json:"name"`
	Date     int       `json:"dt"`
	Info     struct {
		Country string `json:"country"`
	} `json:"sys"`
}
