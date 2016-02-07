package engine

const (
	BLACK_PAWN   = -1
	BLACK_KNIGHT = -2
	BLACK_BISHOP = -3
	BLACK_ROOK   = -4
	BLACK_QUEEN  = -5
	BLACK_KING   = -6
	EMPTY_SQUARE = 0
	WHITE_PAWN   = 1
	WHITE_KNIGHT = 2
	WHITE_BISHOP = 3
	WHITE_ROOK   = 4
	WHITE_QUEEN  = 5
	WHITE_KING   = 6

	OUT_OF_BOARD = 99
)

type Board struct {
	pieces            [8][8]int8
	WhiteKingHasMoved bool
	BlackKingHasMoved bool
}

func NewBoard() *Board {
	b := &Board{
		WhiteKingHasMoved: false,
		BlackKingHasMoved: false,
	}
	return b
}

func (b *Board) findRookMoves(col int8, row int8, c Color) []Move {
	moves := []Move{}

	for j := row + 1; b.isEmptyOrOppositeColor(col, j, c); j = j + 1 {
		moves = append(moves, Move{col, row, col, j})
		if !b.isEmptySquare(col, j) {
			break
		}
	}
	for j := row - 1; b.isEmptyOrOppositeColor(col, j, c); j = j - 1 {
		moves = append(moves, Move{col, row, col, j})
		if !b.isEmptySquare(col, j) {
			break
		}
	}
	for i := col + 1; b.isEmptyOrOppositeColor(i, row, c); i = i + 1 {
		moves = append(moves, Move{col, row, i, row})
		if !b.isEmptySquare(i, row) {
			break
		}
	}
	for i := col - 1; b.isEmptyOrOppositeColor(i, row, c); i = i - 1 {
		moves = append(moves, Move{col, row, i, row})
		if !b.isEmptySquare(i, row) {
			break
		}
	}
	return moves
}

func (b *Board) findBishopMoves(col int8, row int8, c Color) []Move {
	moves := []Move{}

	for i, j := col+1, row+1; b.isEmptyOrOppositeColor(i, j, c); i, j = i+1, j+1 {
		moves = append(moves, Move{col, row, i, j})
		if !b.isEmptySquare(i, j) {
			break
		}
	}
	for i, j := col-1, row+1; b.isEmptyOrOppositeColor(i, j, c); i, j = i-1, j+1 {
		moves = append(moves, Move{col, row, i, j})
		if !b.isEmptySquare(i, j) {
			break
		}
	}
	for i, j := col+1, row-1; b.isEmptyOrOppositeColor(i, j, c); i, j = i+1, j-1 {
		moves = append(moves, Move{col, row, i, j})
		if !b.isEmptySquare(i, j) {
			break
		}
	}
	for i, j := col-1, row-1; b.isEmptyOrOppositeColor(i, j, c); i, j = i-1, j-1 {
		moves = append(moves, Move{col, row, i, j})
		if !b.isEmptySquare(i, j) {
			break
		}
	}
	return moves
}

func (b *Board) findKingMoves(col int8, row int8, c Color) []Move {
	moves := []Move{}

	if b.isEmptyOrOppositeColor(col-1, row-1, c) {
		moves = append(moves, Move{col, row, col - 1, row - 1})
	}
	if b.isEmptyOrOppositeColor(col-1, row, c) {
		moves = append(moves, Move{col, row, col - 1, row})
	}
	if b.isEmptyOrOppositeColor(col-1, row+1, c) {
		moves = append(moves, Move{col, row, col - 1, row + 1})
	}
	if b.isEmptyOrOppositeColor(col, row-1, c) {
		moves = append(moves, Move{col, row, col, row - 1})
	}
	if b.isEmptyOrOppositeColor(col, row+1, c) {
		moves = append(moves, Move{col, row, col, row + 1})
	}
	if b.isEmptyOrOppositeColor(col+1, row-1, c) {
		moves = append(moves, Move{col, row, col + 1, row - 1})
	}
	if b.isEmptyOrOppositeColor(col+1, row, c) {
		moves = append(moves, Move{col, row, col + 1, row})
	}
	if b.isEmptyOrOppositeColor(col+1, row+1, c) {
		moves = append(moves, Move{col, row, col + 1, row + 1})
	}
	if c == WHITE && col == 4 && row == 7 && b.getSquare(7, 7) == WHITE_ROOK && b.isEmptySquare(col+1, row) && b.isEmptySquare(col+2, row) {
		moves = append(moves, Move{col, row, col + 2, row})
	}
	if c == WHITE && col == 4 && row == 7 && b.getSquare(0, 7) == WHITE_ROOK && b.isEmptySquare(col-1, row) && b.isEmptySquare(col-2, row) && b.isEmptySquare(col-3, row) {
		moves = append(moves, Move{col, row, col - 2, row})
	}
	if c == BLACK && col == 4 && row == 0 && b.getSquare(7, 0) == BLACK_ROOK && b.isEmptySquare(col+1, row) && b.isEmptySquare(col+2, row) {
		moves = append(moves, Move{col, row, col + 2, row})
	}
	if c == BLACK && col == 4 && row == 7 && b.getSquare(0, 7) == BLACK_ROOK && b.isEmptySquare(col-1, row) && b.isEmptySquare(col-2, row) && b.isEmptySquare(col-3, row) {
		moves = append(moves, Move{col, row, col - 2, row})
	}
	return moves
}

func (b *Board) findKnightMoves(col int8, row int8, c Color) []Move {
	moves := []Move{}

	if b.isEmptyOrOppositeColor(col-2, row-1, c) {
		moves = append(moves, Move{col, row, col - 2, row - 1})
	}
	if b.isEmptyOrOppositeColor(col+2, row-1, c) {
		moves = append(moves, Move{col, row, col + 2, row - 1})
	}
	if b.isEmptyOrOppositeColor(col-1, row-2, c) {
		moves = append(moves, Move{col, row, col - 1, row - 2})
	}
	if b.isEmptyOrOppositeColor(col+1, row-2, c) {
		moves = append(moves, Move{col, row, col + 1, row - 2})
	}
	if b.isEmptyOrOppositeColor(col-2, row+1, c) {
		moves = append(moves, Move{col, row, col - 2, row + 1})
	}
	if b.isEmptyOrOppositeColor(col+2, row+1, c) {
		moves = append(moves, Move{col, row, col + 2, row + 1})
	}
	if b.isEmptyOrOppositeColor(col-1, row+2, c) {
		moves = append(moves, Move{col, row, col - 1, row + 2})
	}
	if b.isEmptyOrOppositeColor(col+1, row+2, c) {
		moves = append(moves, Move{col, row, col + 1, row + 2})
	}
	return moves
}

func (b *Board) findMovesOfSquare(col int8, row int8, c Color) []Move {
	moves := []Move{}
	if !b.isEmptySquare(col, row) {
		piece := b.getSquare(col, row)
		if c == WHITE {
			if piece == WHITE_PAWN {
				if b.getSquare(col, row-1) == EMPTY_SQUARE {
					moves = append(moves, Move{col, row, col, row - 1})
				}
				if row == 6 && b.getSquare(col, row-1) == EMPTY_SQUARE && b.getSquare(col, row-2) == EMPTY_SQUARE {
					moves = append(moves, Move{col, row, col, row - 2})
				}
				if b.isOppositeColor(col-1, row-1, c) {
					moves = append(moves, Move{col, row, col - 1, row - 1})
				}
				if b.isOppositeColor(col+1, row-1, c) {
					moves = append(moves, Move{col, row, col + 1, row - 1})
				}
			} else if piece == WHITE_BISHOP {
				moves = append(moves, b.findBishopMoves(col, row, c)...)
			} else if piece == WHITE_ROOK {
				moves = append(moves, b.findRookMoves(col, row, c)...)
			} else if piece == WHITE_QUEEN {
				moves = append(moves, b.findBishopMoves(col, row, c)...)
				moves = append(moves, b.findRookMoves(col, row, c)...)
			} else if piece == WHITE_KNIGHT {
				moves = append(moves, b.findKnightMoves(col, row, c)...)
			} else if piece == WHITE_KING {
				moves = append(moves, b.findKingMoves(col, row, c)...)
			}
		} else {
			if piece == BLACK_PAWN {
				if b.getSquare(col, row+1) == EMPTY_SQUARE {
					moves = append(moves, Move{col, row, col, row + 1})
				}
				if row == 1 && b.getSquare(col, row+1) == EMPTY_SQUARE && b.getSquare(col, row+2) == EMPTY_SQUARE {
					moves = append(moves, Move{col, row, col, row + 2})
				}
				if b.isOppositeColor(col-1, row+1, c) {
					moves = append(moves, Move{col, row, col - 1, row + 1})
				}
				if b.isOppositeColor(col+1, row+1, c) {
					moves = append(moves, Move{col, row, col + 1, row + 1})
				}
			} else if piece == BLACK_BISHOP {
				moves = append(moves, b.findBishopMoves(col, row, c)...)
			} else if piece == BLACK_ROOK {
				moves = append(moves, b.findRookMoves(col, row, c)...)
			} else if piece == BLACK_QUEEN {
				moves = append(moves, b.findBishopMoves(col, row, c)...)
				moves = append(moves, b.findRookMoves(col, row, c)...)
			} else if piece == BLACK_KNIGHT {
				moves = append(moves, b.findKnightMoves(col, row, c)...)
			} else if piece == BLACK_KING {
				moves = append(moves, b.findKingMoves(col, row, c)...)
			}
		}
	}
	return moves
}

func (b *Board) findAllMoves(c Color) []Move {
	moves := []Move{}
	for i := int8(0); i < 8; i++ {
		for j := int8(0); j < 8; j++ {
			square_moves := b.findMovesOfSquare(i, j, c)
			if len(square_moves) > 0 {
				moves = append(moves, square_moves...)
			}
		}
	}
	return moves
}

func (b *Board) empty() {
	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			b.pieces[i][j] = EMPTY_SQUARE
		}
	}
}

func (b *Board) Reset() {
	b.empty()
	b.setSquare(0, 7, WHITE_ROOK)
	b.setSquare(1, 7, WHITE_KNIGHT)
	b.setSquare(2, 7, WHITE_BISHOP)
	b.setSquare(3, 7, WHITE_QUEEN)
	b.setSquare(4, 7, WHITE_KING)
	b.setSquare(5, 7, WHITE_BISHOP)
	b.setSquare(6, 7, WHITE_KNIGHT)
	b.setSquare(7, 7, WHITE_ROOK)

	b.setSquare(0, 6, WHITE_PAWN)
	b.setSquare(1, 6, WHITE_PAWN)
	b.setSquare(2, 6, WHITE_PAWN)
	b.setSquare(3, 6, WHITE_PAWN)
	b.setSquare(4, 6, WHITE_PAWN)
	b.setSquare(5, 6, WHITE_PAWN)
	b.setSquare(6, 6, WHITE_PAWN)
	b.setSquare(7, 6, WHITE_PAWN)

	b.setSquare(0, 0, BLACK_ROOK)
	b.setSquare(1, 0, BLACK_KNIGHT)
	b.setSquare(2, 0, BLACK_BISHOP)
	b.setSquare(3, 0, BLACK_QUEEN)
	b.setSquare(4, 0, BLACK_KING)
	b.setSquare(5, 0, BLACK_BISHOP)
	b.setSquare(6, 0, BLACK_KNIGHT)
	b.setSquare(7, 0, BLACK_ROOK)

	b.setSquare(0, 1, BLACK_PAWN)
	b.setSquare(1, 1, BLACK_PAWN)
	b.setSquare(2, 1, BLACK_PAWN)
	b.setSquare(3, 1, BLACK_PAWN)
	b.setSquare(4, 1, BLACK_PAWN)
	b.setSquare(5, 1, BLACK_PAWN)
	b.setSquare(6, 1, BLACK_PAWN)
	b.setSquare(7, 1, BLACK_PAWN)
}

func (b *Board) setSquare(col int8, row int8, piece int8) {
	b.pieces[row][col] = piece
}

func (b *Board) getSquare(col int8, row int8) int8 {
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return OUT_OF_BOARD
	}
	return b.pieces[row][col]
}

func (b *Board) isEmptyOrOppositeColor(col int8, row int8, color Color) bool {
	return b.isEmptySquare(col, row) || b.isOppositeColor(col, row, color)
}

func (b *Board) isOppositeColor(col int8, row int8, color Color) bool {
	if b.getSquare(col, row) == OUT_OF_BOARD {
		return false
	}
	return squareColor(b.getSquare(col, row)) == color.other()
}

func (b *Board) isEmptySquare(col int8, row int8) bool {
	if b.getSquare(col, row) == OUT_OF_BOARD {
		return false
	}
	return b.getSquare(col, row) == EMPTY_SQUARE
}

func (b *Board) MakeMove(m Move) int8 {
	capturedPiece := b.getSquare(m.ToCol, m.ToRow)
	piece := b.getSquare(m.FromCol, m.FromRow)
	b.setSquare(m.FromCol, m.FromRow, EMPTY_SQUARE)
	b.setSquare(m.ToCol, m.ToRow, piece)
	if piece == WHITE_KING && m.FromRow == 7 && m.FromCol == 4 && m.ToCol == 6 {
		b.setSquare(5, 7, WHITE_ROOK)
		b.setSquare(7, 7, EMPTY_SQUARE)
	}
	if piece == WHITE_KING && m.FromRow == 7 && m.FromCol == 4 && m.ToCol == 2 {
		b.setSquare(3, 7, WHITE_ROOK)
		b.setSquare(0, 7, EMPTY_SQUARE)
	}
	if piece == BLACK_KING && m.FromRow == 0 && m.FromCol == 4 && m.ToCol == 6 {
		b.setSquare(5, 0, BLACK_ROOK)
		b.setSquare(7, 0, EMPTY_SQUARE)
	}
	if piece == WHITE_KING && m.FromRow == 0 && m.FromCol == 4 && m.ToCol == 2 {
		b.setSquare(3, 0, BLACK_ROOK)
		b.setSquare(0, 0, EMPTY_SQUARE)
	}
	return capturedPiece
}

func (b *Board) reverseMove(m Move, capturedPiece int8) {
	piece := b.getSquare(m.ToCol, m.ToRow)
	b.setSquare(m.FromCol, m.FromRow, piece)
	b.setSquare(m.ToCol, m.ToRow, capturedPiece)
	if piece == WHITE_KING && m.FromRow == 7 && m.FromCol == 4 && m.ToCol == 6 {
		b.setSquare(7, 7, WHITE_ROOK)
		b.setSquare(5, 7, EMPTY_SQUARE)
	}
	if piece == WHITE_KING && m.FromRow == 7 && m.FromCol == 4 && m.ToCol == 2 {
		b.setSquare(0, 7, WHITE_ROOK)
		b.setSquare(3, 7, EMPTY_SQUARE)
	}
	if piece == BLACK_KING && m.FromRow == 0 && m.FromCol == 4 && m.ToCol == 6 {
		b.setSquare(7, 0, BLACK_ROOK)
		b.setSquare(5, 0, EMPTY_SQUARE)
	}
	if piece == WHITE_KING && m.FromRow == 0 && m.FromCol == 4 && m.ToCol == 2 {
		b.setSquare(0, 0, BLACK_ROOK)
		b.setSquare(3, 0, EMPTY_SQUARE)
	}
}
