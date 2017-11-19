package pluto_test

import (
	"github.com/jamesjoshuahill/pluto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {
	It("has a zero value of 0,0 heading north", func() {
		position := pluto.Position{}
		Expect(position.X).To(Equal(0))
		Expect(position.Y).To(Equal(0))
		Expect(position.Heading).To(Equal(pluto.NORTH))
	})

	It("knows the position one step forward", func() {
		position := pluto.Position{}
		Expect(position.Forward()).To(Equal(pluto.Position{X: 0, Y: 1}))
	})

	It("knows the position one step backward", func() {
		position := pluto.Position{}
		Expect(position.Backward()).To(Equal(pluto.Position{X: 0, Y: -1}))
	})

	It("knows east is right of north", func() {
		position := pluto.Position{}
		Expect(position.Right()).To(Equal(pluto.Position{Heading: pluto.EAST}))
	})

	It("knows west is left of north", func() {
		position := pluto.Position{}
		Expect(position.Left()).To(Equal(pluto.Position{Heading: pluto.WEST}))
	})
})
