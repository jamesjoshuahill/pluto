package pluto

type Heading int

const (
	NORTH Heading = iota
	EAST
	_
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
	p.Heading++
	return p
}

func (p Position) Left() Position {
	p.Heading = WEST
	return p
}
