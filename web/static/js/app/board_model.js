var WHITE_PAWN   = 0;
var WHITE_KNIGHT = 1;
var WHITE_BISHOP = 2;
var WHITE_ROOK   = 3;
var WHITE_QUEEN  = 4;
var WHITE_KING   = 5;

var BLACK_PAWN   = 6;
var BLACK_KNIGHT = 7;
var BLACK_BISHOP = 8;
var BLACK_ROOK   = 9;
var BLACK_QUEEN  = 10;
var BLACK_KING   = 11;

var EMPTY_SQUARE = 12;
var OUT_OF_BOARD = 13;

var PIECE_COLOR_BLACK = -1;
var PIECE_COLOR_NEUTRAL = 0;
var PIECE_COLOR_WHITE = 1;

var InvalidCoordinateError = new Error("Invalid coordinate");
var IllegalMoveError = new Error("Illegal Move");

function BoardModel() {
	this.board = [];
	for (var i = 0; i < 8; i++) {
		var row = [];
		for (var j = 0; j < 8; j++) {
			row.push(EMPTY_SQUARE);
		}
		this.board.push(row);
	}
}

BoardModel.prototype.set_coord = function (row, col, piece) {
	if ((row < 0) || (row > 7) ||
		(col < 0) || (col > 7) ||
		(piece < 0) || (piece > 12)) {
		throw InvalidCoordinateError;
	}
	this.board[row][col] = piece;
}

BoardModel.prototype.get_coord = function (row, col) {
	if ((row < 0) || (row > 7) ||
		(col < 0) || (col > 7)) {
		return OUT_OF_BOARD;
	}
	return this.board[row][col];
}

BoardModel.prototype.set_square = function (square, piece) {
	var col = square.charCodeAt(0) - 97;
	var row = parseInt(square.charAt(1)) - 1;
	this.set_coord(row, col, piece);
}

BoardModel.prototype.get_square = function (square) {
	var col = square.charCodeAt(0) - 97;
	var row = parseInt(square.charAt(1)) - 1;
	return this.get_coord(row, col);
}

BoardModel.prototype.reset_pieces = function () {
	this.set_square("a1", WHITE_ROOK);
	this.set_square("b1", WHITE_KNIGHT);
	this.set_square("c1", WHITE_BISHOP);
	this.set_square("d1", WHITE_QUEEN);
	this.set_square("e1", WHITE_KING);
	this.set_square("f1", WHITE_BISHOP);
	this.set_square("g1", WHITE_KNIGHT);
	this.set_square("h1", WHITE_ROOK);

	this.set_square("a2", WHITE_PAWN);
	this.set_square("b2", WHITE_PAWN);
	this.set_square("c2", WHITE_PAWN);
	this.set_square("d2", WHITE_PAWN);
	this.set_square("e2", WHITE_PAWN);
	this.set_square("f2", WHITE_PAWN);
	this.set_square("g2", WHITE_PAWN);
	this.set_square("h2", WHITE_PAWN);

	this.set_square("a7", BLACK_PAWN);
	this.set_square("b7", BLACK_PAWN);
	this.set_square("c7", BLACK_PAWN);
	this.set_square("d7", BLACK_PAWN);
	this.set_square("e7", BLACK_PAWN);
	this.set_square("f7", BLACK_PAWN);
	this.set_square("g7", BLACK_PAWN);
	this.set_square("h7", BLACK_PAWN);

	this.set_square("a8", BLACK_ROOK);
	this.set_square("b8", BLACK_KNIGHT);
	this.set_square("c8", BLACK_BISHOP);
	this.set_square("d8", BLACK_QUEEN);
	this.set_square("e8", BLACK_KING);
	this.set_square("f8", BLACK_BISHOP);
	this.set_square("g8", BLACK_KNIGHT);
	this.set_square("h8", BLACK_ROOK);
}

BoardModel.prototype.make_move = function (square_from, square_to) {
	var piece = this.get_square(square_from);
	var legal_moves = this.get_moves_of_square(square_from);
	if (!_.include(legal_moves, square_to)) {
		console.log(legal_moves);
		throw IllegalMoveError;
	}
	this.set_square(square_to, piece);
	this.set_square(square_from, EMPTY_SQUARE);
    if (square_from == "e1" && square_to == "g1") {
      this.set_square("f1", WHITE_ROOK);
      this.set_square("h1", EMPTY_SQUARE);
    }
    if (square_from == "e1" && square_to == "c1") {
      this.set_square("d1", WHITE_ROOK);
      this.set_square("a1", EMPTY_SQUARE);
    }
    if (square_from == "e8" && square_to == "g8") {
      this.set_square("f8", BLACK_ROOK);
      this.set_square("h8", EMPTY_SQUARE);
    }
    if (square_from == "e8" && square_to == "c8") {
      this.set_square("d8", BLACK_ROOK);
      this.set_square("a8", EMPTY_SQUARE);
    }
}

BoardModel.prototype.get_piece_color = function (piece) {
	if (piece <= WHITE_KING) {
		return PIECE_COLOR_WHITE;
	} else if (piece <= BLACK_KING) {
		return PIECE_COLOR_BLACK;
	} else if (piece == OUT_OF_BOARD) {
		return OUT_OF_BOARD;
	} else {
		return PIECE_COLOR_NEUTRAL;
	}
}

BoardModel.prototype.get_moves_of_square = function (square) {
	var is_square_movable = function (square_value) {
		var to_piece_color = BoardModel.prototype.get_piece_color.call(this, square_value);
		if (square_value ==  EMPTY_SQUARE) {
			return 1;
		} else if (to_piece_color == -1 * piece_color) {
			return 2;
		} else {
			return 0;
		}
	};
	var coord_to_square = function (row, col) {
		return String.fromCharCode(97 + col) + (row + 1);
	};
	var moves = [];
	var piece = this.get_square(square);
	var piece_color = this.get_piece_color(piece);
	var col = square.charCodeAt(0) - 97;
	var row = parseInt(square.charAt(1)) - 1;

	if (piece == WHITE_PAWN) {
		if (this.get_coord(row + 1, col) == EMPTY_SQUARE) {
			moves.push(coord_to_square(row + 1, col));
		}
		if ((row == 1) &&
			(this.get_coord(row + 1, col) == EMPTY_SQUARE) &&
			(this.get_coord(row + 2, col) == EMPTY_SQUARE)) {
			moves.push(coord_to_square(row + 2, col));
		}
		if (this.get_piece_color(
				this.get_coord(row + 1, col - 1)) == -1 * piece_color) {
			moves.push(coord_to_square(row + 1, col - 1));
		}
		if (this.get_piece_color(
				this.get_coord(row + 1, col + 1)) == -1 * piece_color) {
			moves.push(coord_to_square(row + 1, col + 1));
		}
	}

	if (piece == BLACK_PAWN) {
		if (this.get_coord(row - 1, col) == EMPTY_SQUARE) {
			moves.push(coord_to_square(row - 1, col));
		}
		if ((row == 6) &&
			(this.get_coord(row - 1, col) == EMPTY_SQUARE) &&
			(this.get_coord(row - 2, col) == EMPTY_SQUARE)) {
			moves.push(coord_to_square(row - 2, col));
		}
		if (this.get_piece_color(
				this.get_coord(row - 1, col - 1)) == -1 * piece_color) {
			moves.push(coord_to_square(row - 1, col - 1));
		}
		if (this.get_piece_color(
				this.get_coord(row - 1, col + 1)) == -1 * piece_color) {
			moves.push(coord_to_square(row - 1, col + 1));
		}
	}

	if ((piece == WHITE_KNIGHT) || (piece == BLACK_KNIGHT)) {
		if (is_square_movable(this.get_coord(row + 1, col + 2))) {
			moves.push(coord_to_square(row + 1, col + 2));
		}
		if (is_square_movable(this.get_coord(row + 1, col - 2))) {
			moves.push(coord_to_square(row + 1, col - 2));
		}
		if (is_square_movable(this.get_coord(row - 1, col + 2))) {
			moves.push(coord_to_square(row - 1, col + 2));
		}
		if (is_square_movable(this.get_coord(row - 1, col - 2))) {
			moves.push(coord_to_square(row - 1, col - 2));
		}
		if (is_square_movable(this.get_coord(row + 2, col + 1))) {
			moves.push(coord_to_square(row + 2, col + 1));
		}
		if (is_square_movable(this.get_coord(row + 2, col - 1))) {
			moves.push(coord_to_square(row + 2, col - 1));
		}
		if (is_square_movable(this.get_coord(row - 2, col + 1))) {
			moves.push(coord_to_square(row - 2, col + 1));
		}
		if (is_square_movable(this.get_coord(row - 2, col - 1))) {
			moves.push(coord_to_square(row - 2, col - 1));
		}
	}
	if ((piece == WHITE_BISHOP) || (piece == BLACK_BISHOP)) {
		for (var i = row + 1, j = col + 1; i < 8 && j < 8; i += 1, j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row + 1, j = col - 1; i < 8 && j >= 0; i += 1, j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col + 1; i >= 0 && j < 8; i -= 1, j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col - 1; i >= 0 && j >= 0; i -= 1, j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
	}
	if ((piece == WHITE_ROOK) || (piece == BLACK_ROOK)) {
		for (var i = row + 1, j = col; i < 8; i += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col; i >= 0; i -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row, j = col + 1; j < 8; j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row, j = col - 1; j >= 0; j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
	}
	if ((piece == WHITE_QUEEN) || (piece == BLACK_QUEEN)) {
		for (var i = row + 1, j = col; i < 8; i += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col; i >= 0; i -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row, j = col + 1; j < 8; j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row, j = col - 1; j >= 0; j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row + 1, j = col + 1; i < 8 && j < 8; i += 1, j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row + 1, j = col - 1; i < 8 && j >= 0; i += 1, j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col + 1; i >= 0 && j < 8; i -= 1, j += 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
		for (var i = row - 1, j = col - 1; i >= 0 && j >= 0; i -= 1, j -= 1) {
			var sqrm = is_square_movable(this.get_coord(i, j));
			if (sqrm == 1) {
				moves.push(coord_to_square(i, j));
			} else if (sqrm == 2) {
				moves.push(coord_to_square(i, j));
				break;
			} else {
				break;
			}
		}
	}

	if ((piece == WHITE_KING) || (piece == BLACK_KING)) {
		if (is_square_movable(this.get_coord(row + 1, col + 1))) {
			moves.push(coord_to_square(row + 1, col + 1));
		}
		if (is_square_movable(this.get_coord(row + 1, col))) {
			moves.push(coord_to_square(row + 1, col));
		}
		if (is_square_movable(this.get_coord(row + 1, col - 1))) {
			moves.push(coord_to_square(row + 1, col - 1));
		}
		if (is_square_movable(this.get_coord(row, col - 1))) {
			moves.push(coord_to_square(row, col - 1));
		}
		if (is_square_movable(this.get_coord(row, col + 1))) {
			moves.push(coord_to_square(row, col + 1));
		}
		if (is_square_movable(this.get_coord(row - 1, col + 1))) {
			moves.push(coord_to_square(row - 1, col + 1));
		}
		if (is_square_movable(this.get_coord(row - 1, col))) {
			moves.push(coord_to_square(row - 1, col));
		}
		if (is_square_movable(this.get_coord(row - 1, col - 1))) {
			moves.push(coord_to_square(row - 1, col - 1));
		}
        moves.push(coord_to_square(row, col - 2));
        moves.push(coord_to_square(row, col + 2));
	}

	return moves;
}
