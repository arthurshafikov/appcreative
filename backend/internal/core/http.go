package core

type ErrorBag map[string][]string

type ErrorResponse struct {
	Errors ErrorBag `json:"errors"`
}
