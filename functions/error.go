package functions

type Error struct{
	ErrorText string
}

func NewError(text string) Error {
	return Error{ErrorText: text}
}