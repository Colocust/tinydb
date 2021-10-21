package zerrors

type EncodingError struct {
	text string
}

func (e *EncodingError) Error() string {
	return e.text
}

func NewEncodingError(text string) *EncodingError {
	return &EncodingError{text: text}
}
