package pluto

type Rover struct {
	position Position
}

func NewRover() Rover {
	return Rover{}
}

func (r Rover) Position() Position {
	return r.position
}

func (r *Rover) Do(command string) {
	r.position = r.position.Forward()
}

type Position struct {
	X int
	Y int
	Heading Heading
}

func (p Position) Forward() Position {
	p.Y++
	return p
}

type Heading int

const NORTH Heading = 0