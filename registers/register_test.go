package registers_test

import (
	"github.com/jhtrauntvein/go-logic-gates/registers"
	testutils "github.com/jhtrauntvein/go-logic-gates/test-utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Single bit register tests", Ordered, func() {
   It("isolates output until enabled", func() {
      var err error
      register := registers.NewOneBitRegister()
      output := &testutils.Probe{}
      register.ConnectOutputProbe(output, registers.Dq)
      Expect(register.Evaluate()).To(Equal([]bool{false}))
      _, err = register.SetInput(true, registers.DData)
      Expect(err).ToNot(HaveOccurred())
      Expect(output.Value).To(BeFalse())
      // because read-enable is not asserted, the value won't be latched
      register.PulseClock()
      Expect(output.Value).To(BeFalse())
      _, err = register.SetInput(true, registers.DLoad)
      Expect(err).ToNot(HaveOccurred())
      Expect(output.Value).To(BeFalse())

      // even after the clock is pulsed, the cell output will be false because we haven't set write enable
      register.PulseClock()
      Expect(output.Value).To(BeFalse())
      _, err = register.SetInput(true, registers.DWEnable)
      Expect(err).ToNot(HaveOccurred())
      Expect(output.Value).To(BeTrue())
   })
})
