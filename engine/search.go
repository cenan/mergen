package engine

func (b *Board) Evaluate() (result float32) {
	result = 0
	for i := int8(0); i < 8; i += 1 {
		for j := int8(0); j < 8; j += 1 {
			switch b.getSquare(i, j) {
			case WHITE_PAWN:
				result += 1
			case WHITE_ROOK:
				result += 4.5
			case WHITE_KNIGHT:
				result += 2.8
			case WHITE_BISHOP:
				result += 3.1
			case WHITE_QUEEN:
				result += 9
			case WHITE_KING:
				result += 1000
			case BLACK_PAWN:
				result -= 1
			case BLACK_ROOK:
				result -= 4.5
			case BLACK_KNIGHT:
				result -= 2.8
			case BLACK_BISHOP:
				result -= 3.1
			case BLACK_QUEEN:
				result -= 9
			case BLACK_KING:
				result -= 1000
			}
		}
	}
	return
}

func NegaMax(b *Board, depth int, color Color) (float32, []Move) {
	moves := []Move{}
	bestMoves := []Move{}
	bestMove := Move{}
	score := float32(0)
	if depth == 0 {
		points := b.Evaluate()
		if color == BLACK {
			points = -1 * points
		}
		return points, bestMoves
	}
	max := float32(-10000000.0)
	for _, move := range b.findAllMoves(color) {
		piece := b.MakeMove(move)
		score, moves = NegaMax(b, depth-1, color.other())
		score = -1 * score
		b.reverseMove(move, piece)
		if score > max {
			max = score
			bestMoves = moves
			bestMove = move
		}
	}
	return max, append([]Move{bestMove}, bestMoves...)
}
