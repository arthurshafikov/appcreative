package core

import "fmt"

var (
	CityNotDefined         = "City parameter is required"
	ErrInternalServerError = fmt.Errorf("500 Server Error")
	ErrNotFound            = fmt.Errorf("404 Not Found")
	ErrCityNotFound        = fmt.Errorf("city not found")
)
