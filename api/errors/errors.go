package errors

type ApiError string

func (e ApiError) Error() string {
	return string(e)
}

const (
	ErrCheckRequest    ApiError = "check request"
	ErrSetupRequired   ApiError = "setup required"
	ErrInternalError   ApiError = "internal error"
	ErrNoData          ApiError = "no data was present"
	ErrTooManyEntities ApiError = "too many entities to register"
)
