package zerrors

type ParameterError struct {
	text string
}

func NewParameterError(text string) *ParameterError {
	return &ParameterError{text: text}
}

func (e *ParameterError) Error() string {
	return e.text
}
