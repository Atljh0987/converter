package functions

type Either[L Error, R any] struct {
	left   L
	right  R
	isLeft bool
}

func Left[L Error, R any](val L) Either[L, R] {
	return Either[L, R]{left: val, isLeft: true}
}

func Right[L Error, R any](val R) Either[L, R] {
	return Either[L, R]{right: val, isLeft: false}
}

func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

func (e Either[L, R]) Left() L {
	return e.left
}

func (e Either[L, R]) Right() R {
	return e.right
}
