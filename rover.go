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
	switch command {
	case forward:
		r.position = r.position.Forward()
	case backward:
		r.position = r.position.Backward()
	}
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

func (p Position) Backward() Position {
	p.Y--
	return p
}

type Heading int

const NORTH Heading = 0

const (
	forward = "F"
	backward = "B"
)