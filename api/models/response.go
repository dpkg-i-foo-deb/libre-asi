package models

type Status string

const (
	STATUS_OK     Status = "OK"
	STATUS_DENIED Status = "ACCESS_DENIED"
	STATUS_ERROR  Status = "SERVER_ERROR"
	SETUP_REQUIRED
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
