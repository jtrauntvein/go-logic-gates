package truthtable_test

import (
	truthtable "github.com/jtrauntvein/go-logic-gates/truth-table"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("truth table gates test", func() {
	It("and gate passes all tests", func() {
		test_cases := [][]bool{
			{false, false, false},
			{false, true, false},
			{true, false, false},
			{true, true, true},
		}
		gate := truthtable.NewAndGate2()
		default_value := gate.Evaluate()
		Expect(default_value).To(HaveLen(1))
		Expect(default_value[0]).To(BeFalse())
		for _, test := range test_cases {
			_, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			results, err := gate.SetInput(test[1], 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(results[0]).To(Equal(test[2]))
		}
	})

	It("or gate passes all tests", func() {
		test_cases := [][]bool{
			{false, false, false},
			{false, true, true},
			{true, false, true},
			{true, true, true},
		}
		gate := truthtable.NewOr2Gate()
		default_value := gate.Evaluate()
		Expect(default_value).To(HaveLen(1))
		Expect(default_value[0]).To(BeFalse())
		for _, test := range test_cases {
			_, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			results, err := gate.SetInput(test[1], 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(results[0]).To(Equal(test[2]))
		}
	})

	It("nand gate passes all tests", func() {
		test_cases := [][]bool{
			{false, false, true},
			{false, true, true},
			{true, false, true},
			{true, true, false},
		}
		gate := truthtable.NewNandGate2()
		default_value := gate.Evaluate()
		Expect(default_value).To(HaveLen(1))
		Expect(default_value[0]).To(BeTrue())
		for _, test := range test_cases {
			_, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			results, err := gate.SetInput(test[1], 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(results[0]).To(Equal(test[2]))
		}
	})

	It("nor gate passes all tests", func() {
		test_cases := [][]bool{
			{false, false, true},
			{false, true, false},
			{true, false, false},
			{true, true, false},
		}
		gate := truthtable.NewNorGate2()
		default_value := gate.Evaluate()
		Expect(default_value).To(HaveLen(1))
		Expect(default_value[0]).To(BeTrue())
		for _, test := range test_cases {
			_, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			results, err := gate.SetInput(test[1], 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(results[0]).To(Equal(test[2]))
		}
	})

	It("xor gate passes all tests", func() {
		test_cases := [][]bool{
			{false, false, false},
			{false, true, true},
			{true, false, true},
			{true, true, false},
		}
		gate := truthtable.NewXorGate2()
		default_value := gate.Evaluate()
		Expect(default_value).To(HaveLen(1))
		Expect(default_value[0]).To(BeFalse())
		for _, test := range test_cases {
			_, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			results, err := gate.SetInput(test[1], 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(results[0]).To(Equal(test[2]))
		}
	})

	It("not gate passes all tests", func() {
		test_cases := [][]bool{
			{false, true},
			{true, false},
		}
		gate := truthtable.NewNotGate()
		Expect(gate.Evaluate()).To(Equal([]bool{true}))
		for _, test := range test_cases {
			output, err := gate.SetInput(test[0], 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(HaveLen(1))
			Expect(output[0]).To(Equal(test[1]))
		}
	})
})
