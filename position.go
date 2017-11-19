package pluto

type Heading int

const (
	NORTH Heading = iota
	EAST
	SOUTH
	WEST
)

const gridWidth = 100

type Position struct {
	X       int
	Y       int
	Heading Heading
}

func (p Position) Forward() Position {
	switch p.Heading {
	case NORTH:
		p.Y++
	case EAST:
		p.X++
	case SOUTH:
		p.Y--
	case WEST:
		p.X--
	}
	return p
}

func (p Position) Backward() Position {
	switch p.Heading {
	case NORTH:
		if p.Y == 0 {
			p.Y = gridWidth
		}
		p.Y--

	case EAST:
		if p.X == 0 {
			p.X = gridWidth
		}
		p.X--
	case SOUTH:
		p.Y++
	case WEST:
		p.X++
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
