package main

import "fmt"

type piece int

const (
	pawn piece = iota
	knight
	bishop
	rook
	queen
	king
)

func (p piece) String() string {
	switch p {
	case pawn:
		return "Pawn"
	case knight:
		return "Knight"
	case bishop:
		return "Bishop"
	case rook:
		return "Rook"
	case queen:
		return "Queen"
	case king:
		return "King"
	default:
		return fmt.Sprintf("unknown piece %d", p)
	}
}

var allPieces = []piece{pawn, knight, bishop, rook, queen, king}

type room struct {
	name  string
	piece piece
}

var allRooms = []room{
	{name: "Attic", piece: bishop},
	{name: "Bedroom", piece: pawn},
	{name: "Chapel", piece: bishop},
	{name: "Clocktower", piece: rook},
	{name: "Conservatory", piece: rook},
	{name: "Den", piece: pawn},
	{name: "Dining hall", piece: pawn},
	{name: "Dormitory", piece: pawn},
	{name: "Drafting room", piece: pawn},
	{name: "Guest Room", piece: pawn},
	{name: "Her Ladyship's Den", piece: queen},
	{name: "Nursery", piece: pawn},
	{name: "Observatory", piece: knight},
	{name: "Office", piece: king},
	{name: "Parlor", piece: pawn},
	{name: "Rumpus Room", piece: bishop},
	{name: "Security", piece: knight},
	{name: "Solarium", piece: pawn},
	{name: "Store Room", piece: pawn},
	{name: "Study", piece: queen},
	{name: "Throne Room", piece: king},
	{name: "Vault", piece: rook},
}

type model struct {
	selected  map[int]struct{}
	cursor    int
	collected map[piece]struct{}
	needed    map[piece]struct{}
}
