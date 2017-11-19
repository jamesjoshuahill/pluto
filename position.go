package pluto

type Heading int

const (
	NORTH Heading = 0
	EAST  Heading = 90
	SOUTH Heading = 180
	WEST  Heading = 270

	gridSize = 100

	maxDegrees  = 360
	turnDegrees = 90
)

type Position struct {
	X       int
	Y       int
	Heading Heading
}

func (p Position) Forward() Position {
	switch p.Heading {
	case NORTH:
		p.Y = increment(p.Y)
	case EAST:
		p.X = increment(p.X)
	case SOUTH:
		p.Y = decrement(p.Y)
	case WEST:
		p.X = decrement(p.X)
	}
	return p
}

func (p Position) Backward() Position {
	switch p.Heading {
	case NORTH:
		p.Y = decrement(p.Y)
	case EAST:
		p.X = decrement(p.X)
	case SOUTH:
		p.Y = increment(p.Y)
	case WEST:
		p.X = increment(p.X)
	}
	return p
}

func (p Position) Right() Position {
	p.Heading = (p.Heading + turnDegrees) % maxDegrees
	return p
}

func (p Position) Left() Position {
	p.Heading = (maxDegrees + p.Heading - turnDegrees) % maxDegrees
	return p
}

func increment(i int) int {
	return (i + 1) % gridSize
}

func decrement(i int) int {
	return (gridSize + i - 1) % gridSize
}
