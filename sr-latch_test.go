package gologicgates_test

import (
	gologicgates "github.com/jhtrauntvein/go-logic-gates"
	"github.com/jhtrauntvein/go-logic-gates/latches"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type latch_probe struct {
	value bool
	count int
}

func (p *latch_probe) OnLineChanged(line *gologicgates.Line, value bool) {
	p.value = value
	p.count++
}

var _ = Describe("S-R Latch Unit Tests", func() {
	It("can construct an sr-latch and default values are as expected", func() {
		var err error
		latch := latches.NewSrLatch()
		q_probe := &latch_probe{}
		qn_probe := &latch_probe{}
		latch.ConnectOutputProbe(q_probe, latches.SrQ, true)
		latch.ConnectOutputProbe(qn_probe, latches.SrQNot, true)
		Expect(q_probe.value).To(BeTrue())
		Expect(qn_probe.value).To(BeFalse())
		Expect(q_probe.count).To(Equal(1))
		Expect(qn_probe.count).To(Equal(1))
		_, err = latch.SetInput(true, latches.SrSet)
		Expect(err).ToNot(HaveOccurred())
		Expect(q_probe.value).To(BeFalse())
		Expect(qn_probe.value).To(BeTrue())
		_, err = latch.SetInput(false, latches.SrSet)
		Expect(err).ToNot(HaveOccurred())
		Expect(q_probe.value).To(BeFalse())
		Expect(qn_probe.value).To(BeTrue())
		_, err = latch.SetInput(true, latches.SrReset)
		Expect(err).ToNot(HaveOccurred())
		Expect(q_probe.value).To(BeTrue())
		Expect(qn_probe.value).To(BeFalse())
		_, err = latch.SetInput(false, latches.SrReset)
		Expect(err).ToNot(HaveOccurred())
		Expect(q_probe.value).To(BeTrue())
	})
})
