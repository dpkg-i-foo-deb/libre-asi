package models

type Status string

const (
	STATUS_OK     Status = "OK"
	STATUS_DENIED Status = "ACCESS_DENIED"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
