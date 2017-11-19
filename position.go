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
		p.incrementY()
	case EAST:
		p.incrementX()
	case SOUTH:
		p.decrementY()
	case WEST:
		p.decrementX()
	}
	return p
}

func (p Position) Backward() Position {
	switch p.Heading {
	case NORTH:
		p.decrementY()
	case EAST:
		p.decrementX()
	case SOUTH:
		p.incrementY()
	case WEST:
		p.incrementX()
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

func (p *Position) incrementX() {
	p.X++
	if p.X == gridSize {
		p.X = 0
	}
}

func (p *Position) incrementY() {
	p.Y++
	if p.Y == gridSize {
		p.Y = 0
	}
}

func (p *Position) decrementX() {
	if p.X == 0 {
		p.X = gridSize
	}
	p.X--

}

func (p *Position) decrementY() {
	if p.Y == 0 {
		p.Y = gridSize
	}
	p.Y--

}
