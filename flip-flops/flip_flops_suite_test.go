package flipflops_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFlipFlops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlipFlops Suite")
}
