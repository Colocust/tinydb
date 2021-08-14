package errors

type TypeError struct {
}

func (e *TypeError) Error() string {
	return "errors type"
}

func NewTypeError() *TypeError {
	return &TypeError{}
}
