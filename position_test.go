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
		Expect(position.X).To(Equal(0))
		Expect(position.Y).To(Equal(0))
		Expect(position.Heading).To(Equal(pluto.NORTH))
	})

	DescribeTable("moving forward",
		func(heading pluto.Heading, expectedX, expectedY int) {
			position := pluto.Position{X: 1, Y: 1, Heading: heading}
			Expect(position.Forward()).To(Equal(pluto.Position{
				X:       expectedX,
				Y:       expectedY,
				Heading: heading,
			}))
		},
		Entry("north", pluto.NORTH, 1, 2),
		Entry("east", pluto.EAST, 2, 1),
		Entry("south", pluto.SOUTH, 1, 0),
		Entry("west", pluto.WEST, 0, 1),
	)

	It("circles the pole moving forwards", func() {
		position := pluto.Position{X: 1, Y: 1, Heading: pluto.WEST}
		position = position.Forward().Forward()
		Expect(position).To(Equal(pluto.Position{X: 99, Y: 1, Heading: pluto.WEST}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{X: 99, Y: 99, Heading: pluto.SOUTH}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{X: 1, Y: 99, Heading: pluto.EAST}))
		position = position.Left().Forward().Forward()
		Expect(position).To(Equal(pluto.Position{X: 1, Y: 1, Heading: pluto.NORTH}))
	})

	DescribeTable("moving backward",
		func(heading pluto.Heading, expectedX, expectedY int) {
			position := pluto.Position{X: 1, Y: 1, Heading: heading}
			Expect(position.Backward()).To(Equal(pluto.Position{
				X:       expectedX,
				Y:       expectedY,
				Heading: heading,
			}))
		},
		Entry("north", pluto.NORTH, 1, 0),
		Entry("east", pluto.EAST, 0, 1),
		Entry("south", pluto.SOUTH, 1, 2),
		Entry("west", pluto.WEST, 2, 1),
	)

	DescribeTable("turning right",
		func(initial, expected pluto.Heading) {
			position := pluto.Position{Heading: initial}
			Expect(position.Right()).To(Equal(pluto.Position{Heading: expected}))
		},
		Entry("north -> east", pluto.NORTH, pluto.EAST),
		Entry("east -> south", pluto.EAST, pluto.SOUTH),
		Entry("south -> west", pluto.SOUTH, pluto.WEST),
		Entry("west -> north", pluto.WEST, pluto.NORTH),
	)

	DescribeTable("turning left",
		func(initial, expected pluto.Heading) {
			position := pluto.Position{Heading: initial}
			Expect(position.Left()).To(Equal(pluto.Position{Heading: expected}))
		},
		Entry("north -> west", pluto.NORTH, pluto.WEST),
		Entry("west -> south", pluto.WEST, pluto.SOUTH),
		Entry("south -> east", pluto.SOUTH, pluto.EAST),
		Entry("east -> north", pluto.EAST, pluto.NORTH),
	)
})
