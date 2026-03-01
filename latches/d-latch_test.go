package latches_test

import (
	"github.com/jhtrauntvein/go-logic-gates/latches"
	testutils "github.com/jhtrauntvein/go-logic-gates/test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("D-Latch unit tests", func() {
	It("can set up and monitor a D-Latch device", func() {
		var err error
		latch := latches.NewDLatch()
		qp := &testutils.Probe{}
		qnp := &testutils.Probe{}
		err = latch.ConnectOutputProbe(qp, latches.SrQ, true)
		Expect(err).ToNot(HaveOccurred())
		err = latch.ConnectOutputProbe(qnp, latches.SrQNot, true)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeFalse())
		Expect(qnp.Value).To(BeTrue())
		_, err = latch.SetInput(true, latches.DData)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeFalse())
		_, err = latch.SetInput(true, latches.DEnable)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeTrue())
		_, err = latch.SetInput(false, latches.DEnable)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeTrue())
		_, err = latch.SetInput(false, latches.DData)
		Expect(qp.Value).To(BeTrue())
		_, err = latch.SetInput(true, latches.DEnable)
		Expect(qp.Value).To(BeFalse())
		_, err = latch.SetInput(false, latches.DEnable)
		Expect(qp.Value).To(BeFalse())
	})
})
