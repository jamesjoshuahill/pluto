package pluto

const (
	NORTH Heading = 0
	EAST  Heading = 90
	SOUTH Heading = 180
	WEST  Heading = 270

	maxDegrees = 360
)

type Heading int

func (h Heading) Turn(degrees Heading) Heading {
	return (maxDegrees + h + degrees) % maxDegrees
}
