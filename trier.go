package trier

type Try[T any] struct {
	Result T     `json:"result"`
	Error  error `json:"error"`
}

func (t Try[T]) Failed() bool {
	return t.Error != nil
}

func (t Try[T]) Succeded() bool {
	return t.Error == nil
}
