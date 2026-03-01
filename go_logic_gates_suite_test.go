package gologicgates_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoLogicGates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoLogicGates Suite")
}
