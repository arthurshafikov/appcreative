package core

import "fmt"

var (
	ErrCityNotDefined = fmt.Errorf("city parameter is required")
	ErrInternalServer = fmt.Errorf("500 Server Error")
	ErrNotFound       = fmt.Errorf("404 Not Found")
	ErrCityNotFound   = fmt.Errorf("city not found")
)
