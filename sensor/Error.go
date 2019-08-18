package sensor

type Error struct {
	message string
}

func NewError(message string) *Error {
	return &Error{message: message}
}

func (error *Error) Error() string {
	return error.message
}