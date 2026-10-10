package models

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count,omitempty"`
}
