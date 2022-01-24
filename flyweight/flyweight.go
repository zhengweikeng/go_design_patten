/**
 * @Author: Seedzheng
 * 享元模式
 */

package flyweight

type Color int

var Red Color = 0
var Black Color = 1

type ChessPieceUint struct {
	id    int
	text  string
	color Color
}

func NewChessPieceUnit(id int, text string, c Color) *ChessPieceUint {
	return &ChessPieceUint{
		id:    id,
		text:  text,
		color: c,
	}
}

var chessPieceUints = map[int]*ChessPieceUint{}

func init() {
	chessPieceUints[1] = NewChessPieceUnit(1, "车", Red)
	chessPieceUints[2] = NewChessPieceUnit(2, "马", Red)
	// ... 以此类推
}

type ChessPiece struct {
	*ChessPieceUint
	PosX int
	PoxY int
}

func NewChessPiece(unit *ChessPieceUint, poxX, poxY int) *ChessPiece {
	return &ChessPiece{
		ChessPieceUint: unit,
		PosX:           poxX,
		PoxY:           poxY,
	}
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	b := ChessBoard{}
	chessPieces := map[int]*ChessPiece{}
	chessPieces[1] = NewChessPiece(chessPieceUints[1], 0, 0)
	chessPieces[1] = NewChessPiece(chessPieceUints[2], 1, 0)
	// ... 以此类推

	b.chessPieces = chessPieces
	return &b
}
