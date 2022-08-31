package core

import "fmt"

var (
	CityNotDefined         = "City parameter is required"
	ErrInternalServerError = fmt.Errorf("500 Server Error")
	ErrCityNotFound        = fmt.Errorf("city not found")
)
