package engine

import "fmt"

type Move struct {
	FromCol int8
	FromRow int8
	ToCol   int8
	ToRow   int8
}

func (m Move) String() string {
	cols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return fmt.Sprintf("%s%d %s%d", cols[m.FromCol], 8-m.FromRow, cols[m.ToCol], 8-m.ToRow)
}
