package registers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRegisters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registers Suite")
}
