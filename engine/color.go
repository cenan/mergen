package engine

type Color int8

func (c Color) other() Color {
	return c * -1
}

const (
	WHITE   = -1
	NEUTRAL = 0
	BLACK   = 1
)

func squareColor(square int8) Color {
	if square < EMPTY_SQUARE {
		return BLACK
	} else if square > EMPTY_SQUARE {
		return WHITE
	} else {
		return NEUTRAL
	}
}
