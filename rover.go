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
	case right:
		r.position = r.position.Right()
	case left:
		r.position = r.position.Left()
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

func (p Position) Right() Position {
	p.Heading++
	return p
}

func (p Position) Left() Position {
	p.Heading = WEST
	return p
}

type Heading int

const (
	NORTH Heading = iota
	EAST
	_
	WEST
)

const (
	forward = "F"
	backward = "B"
	right = "R"
	left = "L"
)