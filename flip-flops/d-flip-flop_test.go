package flipflops_test

import (
	flipflops "github.com/jtrauntvein/go-logic-gates/flip-flops"
	testutils "github.com/jtrauntvein/go-logic-gates/test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("D Flip Flop unit tests", Ordered, func() {
	// create the device and set monitors
	var err error
	var flip_flop *flipflops.DFlipFlop
	qp := &testutils.Probe{}
	qnp := &testutils.Probe{}

	BeforeAll(func() {
		flip_flop = flipflops.NewDFlipFlop()
		err = flip_flop.ConnectOutputProbe(qp, flipflops.Dq, true)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeFalse())
		err = flip_flop.ConnectOutputProbe(qnp, flipflops.DqNot, true)
		Expect(qnp.Value).To(BeTrue())
	})

	BeforeEach(func() {
		// we'll latch in a false value before each to ensure a known state
		_, err := flip_flop.SetInput(false, flipflops.DData)
		Expect(err).ToNot(HaveOccurred())
		flip_flop.PulseClock()
	})

	It("preserves state when data line is changed but not the clock", func() {
		_, err := flip_flop.SetInput(true, flipflops.DData)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeFalse())
		Expect(qnp.Value).To(BeTrue())
	})

	It("latches input data true when the clock is pulsed", func() {
		_, err := flip_flop.SetInput(true, flipflops.DData)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeFalse())
		flip_flop.PulseClock()
		Expect(qp.Value).To(BeTrue())
		_, err = flip_flop.SetInput(false, flipflops.DData)
		Expect(err).ToNot(HaveOccurred())
		Expect(qp.Value).To(BeTrue())
	})
})
