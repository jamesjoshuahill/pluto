package pluto

type Rover struct {}

func NewRover() Rover {
	return Rover{}
}

func (Rover) Position() Position {
	return Position{}
}

type Position struct {
	X int
	Y int
	Heading Heading
}

type Heading int

const NORTH Heading = 0