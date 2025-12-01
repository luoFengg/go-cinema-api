package exceptions

type ConflictError struct {
	ErrorMessage string
}

func NewConflictError(error string) ConflictError {
	return ConflictError{ErrorMessage: error}
}

func (e ConflictError) Error() string {
	return e.ErrorMessage
}
