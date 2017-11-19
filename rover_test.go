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

	It("can move forward", func() {
		rover := pluto.NewRover()
		rover.Do("F")
		Expect(rover.Position()).To(Equal(pluto.Position{X: 0, Y: 1, Heading: pluto.NORTH}))
	})
})
