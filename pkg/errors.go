package pkg

type GenericError struct {
	Code    ErrorCode
	Message string
}

type CustomError struct {
	GenericError
}

func NewCustomError(code ErrorCode, msg string) *CustomError {
	return &CustomError{GenericError{
		Code:    code,
		Message: msg,
	}}
}

func (e *CustomError) Error() string {
	return e.Message
}

type ErrorCode uint8

const (
	PreconditionFailed ErrorCode = iota
)
