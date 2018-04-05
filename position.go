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

type Coordinate struct {
	X int
	Y int
}

type Position struct {
	Coordinate Coordinate
	Heading    Heading
}

func (p Position) Forward() Position {
	switch p.Heading {
	case NORTH:
		p.Coordinate.Y = increment(p.Coordinate.Y)
	case EAST:
		p.Coordinate.X = increment(p.Coordinate.X)
	case SOUTH:
		p.Coordinate.Y = decrement(p.Coordinate.Y)
	case WEST:
		p.Coordinate.X = decrement(p.Coordinate.X)
	}
	return p
}

func (p Position) Backward() Position {
	switch p.Heading {
	case NORTH:
		p.Coordinate.Y = decrement(p.Coordinate.Y)
	case EAST:
		p.Coordinate.X = decrement(p.Coordinate.X)
	case SOUTH:
		p.Coordinate.Y = increment(p.Coordinate.Y)
	case WEST:
		p.Coordinate.X = increment(p.Coordinate.X)
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
