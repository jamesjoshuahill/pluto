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
		p.Y++
		if p.Y == gridSize {
			p.Y = 0
		}
	case EAST:
		p.X++
		if p.X == gridSize {
			p.X = 0
		}
	case SOUTH:
		if p.Y == 0 {
			p.Y = gridSize
		}
		p.Y--
	case WEST:
		if p.X == 0 {
			p.X = gridSize
		}
		p.X--
	}
	return p
}

func (p Position) Backward() Position {
	switch p.Heading {
	case NORTH:
		if p.Y == 0 {
			p.Y = gridSize
		}
		p.Y--
	case EAST:
		if p.X == 0 {
			p.X = gridSize
		}
		p.X--
	case SOUTH:
		p.Y++
		if p.Y == gridSize {
			p.Y = 0
		}
	case WEST:
		p.X++
		if p.X == gridSize {
			p.X = 0
		}
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
