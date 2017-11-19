package pluto_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPluto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pluto Suite")
}
