package pluto

type Heading int

const (
	NORTH Heading = iota
	EAST
	SOUTH
	WEST
)

type Position struct {
	X       int
	Y       int
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
	if p.Heading == WEST {
		p.Heading = NORTH
	} else {
		p.Heading++
	}
	return p
}

func (p Position) Left() Position {
	if p.Heading == NORTH {
		p.Heading = WEST
	} else {
		p.Heading--
	}
	return p
}
