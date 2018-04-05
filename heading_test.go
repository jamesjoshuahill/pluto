package pluto_test

import (
	"github.com/jamesjoshuahill/pluto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heading", func() {
	It("knows how to turn 90 degrees clockwise", func() {
		Expect(pluto.NORTH.Turn(90)).To(Equal(pluto.EAST))
		Expect(pluto.EAST.Turn(90)).To(Equal(pluto.SOUTH))
		Expect(pluto.SOUTH.Turn(90)).To(Equal(pluto.WEST))
		Expect(pluto.WEST.Turn(90)).To(Equal(pluto.NORTH))
	})

	It("knows how to turn 90 degrees anti-clockwise", func() {
		Expect(pluto.NORTH.Turn(-90)).To(Equal(pluto.WEST))
		Expect(pluto.WEST.Turn(-90)).To(Equal(pluto.SOUTH))
		Expect(pluto.SOUTH.Turn(-90)).To(Equal(pluto.EAST))
		Expect(pluto.EAST.Turn(-90)).To(Equal(pluto.NORTH))
	})
})
