package errors

import "errors"

const (
	DB_ERROR   = 600
	GO_ERROR   = 700
	USER_ERROR = 800
)

type SensError struct {
	error
	code int
}

func (this *SensError) Error() string {
	return this.error.Error()
}

func New(code int, message string) *SensError {
	return &SensError{errors.New(message), code}
}

func FromError(code int, err error) *SensError {
	return &SensError{err, code}
}

func GetErrorCode(err *SensError) int {
	return err.code
}
