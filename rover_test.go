package pluto_test

import (
	"github.com/jamesjoshuahill/pluto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rover", func() {
	It("starts at position 0,0 heading north", func() {
		rover := pluto.NewRover()
		Expect(rover.Position()).To(Equal(pluto.Position{X: 0, Y: 0, Heading: pluto.NORTH}))
	})

	It("can do one command", func() {
		rover := pluto.NewRover()
		rover.Do("F")
		Expect(rover.Position()).To(Equal(pluto.Position{X: 0, Y: 1, Heading: pluto.NORTH}))
	})

	It("can do all four commands", func() {
		rover := pluto.NewRover()
		rover.Do("FLBR")
		Expect(rover.Position()).To(Equal(pluto.Position{X: 1, Y: 1, Heading: pluto.NORTH}))
	})

	It("can do many commands", func() {
		rover := pluto.NewRover()
		rover.Do("FFRFF")
		Expect(rover.Position()).To(Equal(pluto.Position{X: 2, Y: 2, Heading: pluto.EAST}))
	})
})
