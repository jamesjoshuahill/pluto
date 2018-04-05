package pluto_test

import (
	"github.com/jamesjoshuahill/pluto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {
	It("has a zero value of 0,0 heading north", func() {
		position := pluto.Position{}
		Expect(position.Coordinate.X).To(Equal(0))
		Expect(position.Coordinate.Y).To(Equal(0))
		Expect(position.Heading).To(Equal(pluto.NORTH))
	})

	DescribeTable("moving forward",
		func(heading pluto.Heading, expectedX, expectedY int) {
			position := pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: heading}
			Expect(position.Forward()).To(Equal(pluto.Position{
				Coordinate: pluto.Coordinate{
					X: expectedX,
					Y: expectedY,
				},
				Heading: heading,
			}))
		},
		Entry("north", pluto.NORTH, 1, 2),
		Entry("east", pluto.EAST, 2, 1),
		Entry("south", pluto.SOUTH, 1, 0),
		Entry("west", pluto.WEST, 0, 1),
	)

	It("circles the pole moving forwards", func() {
		position := pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: pluto.WEST}
		position = position.Forward().Forward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 99, Y: 1}, Heading: pluto.WEST}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 99, Y: 99}, Heading: pluto.SOUTH}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 99}, Heading: pluto.EAST}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: pluto.NORTH}))
	})

	DescribeTable("moving backward",
		func(heading pluto.Heading, expectedX, expectedY int) {
			position := pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: heading}
			Expect(position.Backward()).To(Equal(pluto.Position{
				Coordinate: pluto.Coordinate{
					X: expectedX,
					Y: expectedY,
				},
				Heading: heading,
			}))
		},
		Entry("north", pluto.NORTH, 1, 0),
		Entry("east", pluto.EAST, 0, 1),
		Entry("south", pluto.SOUTH, 1, 2),
		Entry("west", pluto.WEST, 2, 1),
	)

	It("circles the pole moving backwards", func() {
		position := pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: pluto.NORTH}
		position = position.Backward().Backward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 99}, Heading: pluto.NORTH}))
		position = position.Right().Backward().Backward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 99, Y: 99}, Heading: pluto.EAST}))
		position = position.Right().Backward().Backward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 99, Y: 1}, Heading: pluto.SOUTH}))
		position = position.Right().Backward().Backward()
		Expect(position).To(Equal(pluto.Position{Coordinate: pluto.Coordinate{X: 1, Y: 1}, Heading: pluto.WEST}))
	})

	It("turns right", func() {
		position := pluto.Position{Heading: pluto.NORTH}
		Expect(position.Right()).To(Equal(pluto.Position{Heading: pluto.EAST}))
	})

	It("turns left", func() {
		position := pluto.Position{Heading: pluto.NORTH}
		Expect(position.Left()).To(Equal(pluto.Position{Heading: pluto.WEST}))
	})
})
