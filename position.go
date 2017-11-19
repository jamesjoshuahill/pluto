package pluto

type Heading int

const (
	NORTH Heading = iota
	EAST
	SOUTH
	WEST
)

const gridSize = 100

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

func increment(i int) int {
	i++
	if i == gridSize {
		i = 0
	}
	return i
}

func decrement(i int) int {
	if i == 0 {
		i = gridSize
	}
	i--
	return i
}
