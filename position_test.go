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
			position := pluto.Position{Heading: heading}
			Expect(position.Forward()).To(Equal(pluto.Position{X: expectedX, Y: expectedY, Heading: heading}))
		},
		Entry("north", pluto.NORTH, 0, 1),
		Entry("east", pluto.EAST, 1, 0),
		Entry("south", pluto.SOUTH, 0, -1),
		Entry("west", pluto.WEST, -1, 0),
	)

	It("knows the position one step backward", func() {
		position := pluto.Position{}
		Expect(position.Backward()).To(Equal(pluto.Position{X: 0, Y: -1}))
	})

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
