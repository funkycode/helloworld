package errors

type EmptyRequest struct{}

func (e *EmptyRequest) Error() string {
	return "Empty request is not allowed"
}

type EmptyNameNotAllowed struct{}

func (e *EmptyNameNotAllowed) Error() string {
	return "Empty name is not allowed"
}

type ParseError struct{}

func (e *ParseError) Error() string {
	return "Failed to parse request"
}
