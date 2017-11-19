package pluto

const (
	forward  = "F"
	backward = "B"
	right    = "R"
	left     = "L"
)

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
