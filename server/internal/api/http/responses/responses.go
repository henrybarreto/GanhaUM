package responses

type Error struct {
	Error string `json:"error"`
}

func NewError(err error) Error {
	return Error{
		Error: err.Error(),
	}
}
