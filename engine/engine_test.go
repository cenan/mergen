package engine

import (
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	b := NewBoard()
	b.Reset()
	if b.isEmptySquare(0, 5) == false {
		t.Error("a3 square is not empty")
	}
	if b.getSquare(4, 7) != WHITE_KING {
		t.Error("White king is not on e1")
	}

	b.Reset()
	if move_count := len(b.findAllMoves(WHITE)); move_count != 20 {
		for _, x := range b.findAllMoves(WHITE) {
			fmt.Println(x)
		}
		t.Error(fmt.Sprintf("Initial white moves: %d", move_count))
	}

	b.Reset()
	if move_count := len(b.findAllMoves(BLACK)); move_count != 20 {
		for _, x := range b.findAllMoves(BLACK) {
			fmt.Println(x)
		}
		t.Error(fmt.Sprintf("Initial black moves: %d", move_count))
	}

	b.empty()
	b.setSquare(4, 4, WHITE_BISHOP)
	if bishop_move_count := len(b.findAllMoves(WHITE)); bishop_move_count != 13 {
		t.Error(fmt.Sprintf("White Bishop moves: %d", bishop_move_count))
	}
	if bishop_move_count := len(b.findAllMoves(BLACK)); bishop_move_count != 0 {
		t.Error(fmt.Sprintf("Black Bishop moves: %d", bishop_move_count))
	}

	b.empty()
	b.setSquare(4, 4, WHITE_ROOK)
	if rook_move_count := len(b.findAllMoves(WHITE)); rook_move_count != 14 {
		t.Error(fmt.Sprintf("White Rook moves: %d", rook_move_count))
	}

	b.empty()
	b.setSquare(4, 4, WHITE_QUEEN)
	if queen_move_count := len(b.findAllMoves(WHITE)); queen_move_count != 27 {
		t.Error(fmt.Sprintf("White Queen moves: %d", queen_move_count))
	}
	b.empty()
	b.setSquare(4, 4, BLACK_QUEEN)
	if queen_move_count := len(b.findAllMoves(BLACK)); queen_move_count != 27 {
		for _, x := range b.findAllMoves(BLACK) {
			fmt.Println(x)
		}
		t.Error(fmt.Sprintf("Black Queen moves: %d", queen_move_count))
	}

	b.empty()
	b.setSquare(4, 4, WHITE_KING)
	if king_move_count := len(b.findAllMoves(WHITE)); king_move_count != 8 {
		for _, x := range b.findAllMoves(WHITE) {
			fmt.Println(x)
		}
		t.Error(fmt.Sprintf("White King moves: %d", king_move_count))
	}

	b.empty()
	b.setSquare(4, 4, WHITE_KNIGHT)
	if move_count := len(b.findAllMoves(WHITE)); move_count != 8 {
		t.Error(fmt.Sprintf("White knight moves: %d", move_count))
	}

	b.Reset()
	_, moves := NegaMax(b, 5, WHITE)
	fmt.Println(moves)

	b.empty()
	b.setSquare(0, 7, WHITE_ROOK)
	b.setSquare(4, 0, BLACK_KING)
	b.setSquare(3, 1, BLACK_PAWN)
	b.setSquare(4, 1, BLACK_PAWN)
	b.setSquare(5, 1, BLACK_PAWN)
	_, moves = NegaMax(b, 4, WHITE)
	fmt.Println(moves)
}
