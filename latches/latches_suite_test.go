package latches_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLatches(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Latches Suite")
}
