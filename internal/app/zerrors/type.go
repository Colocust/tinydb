package zerrors

type TypeError struct {
	text string
}

func (e *TypeError) Error() string {
	return e.text
}

func NewTypeError(text string) *TypeError {
	return &TypeError{text: text}
}
