package wraped_error

type Error struct {
	Err        error
	StatusCode int
}

func WrapError(err error, statusCode int) *Error {
	return &Error{
		Err:        err,
		StatusCode: statusCode,
	}
}
