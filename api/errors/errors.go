package errors

type ApiError string

func (e ApiError) Error() string {
	return string(e)
}

const (
	ErrCheckRequest        ApiError = "check request"
	ErrSetupRequired       ApiError = "setup required"
	ErrInternalError       ApiError = "internal error"
	ErrNoData              ApiError = "no data was present"
	ErrTooManyEntities     ApiError = "too many entities to register"
	ErrNotImplemmented     ApiError = "functionality not implemmented"
	ErrBadRoute            ApiError = "bad route requested"
	ErrAccessDenied        ApiError = "access denied"
	ErrConflict            ApiError = "data conflict"
	ErrInvalidCredentials  ApiError = "invalid credentials"
	ErrrNeedsPasswordReset ApiError = "password reset needed"
)
