package models

type Status string

const (
	OK             Status = "OK"
	DENIED         Status = "ACCESS_DENIED"
	ERROR          Status = "SERVER_ERROR"
	SETUP_REQUIRED Status = "SETUP_REQUIRED"
	CHECK_REQUEST  Status = "CHECK_REQUEST"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
