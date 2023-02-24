package errors

type ApiError string

func (e ApiError) Error() string {
	return string(e)
}

const (
	ErrCheckRequest  ApiError = "check request"
	ErrSetupRequired ApiError = "setup required"
)
