package core

type WeatherResponse struct {
	City             string           `json:"city"`
	Temperature      int              `json:"temperature"`
	WeatherCondition WeatherCondition `json:"weatherCondition"`
	Wind             WindStatus       `json:"wind"`
}

type WeatherCondition struct {
	Type     string `json:"type"`
	Pressure int    `json:"pressure"`
	Humidity uint8  `json:"humidity"`
}

type WindStatus struct {
	Speed     float64 `json:"speed"`
	Direction string  `json:"direction"`
}
