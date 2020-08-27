package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("just o", func() {
		Expect(Parse("ooo")).To(Equal([]int{0, 0, 0}))
	})
	It("o&i", func() {
		Expect(Parse("ioioio")).To(Equal([]int{1, 2, 3}))
	})
	It("o&i&d", func() {
		Expect(Parse("idoiido")).To(Equal([]int{0, 1}))
	})
	It("o&i&d&s", func() {
		Expect(Parse("isoisoiso")).To(Equal([]int{1, 4, 25}))
	})
	It("codewars", func() {
		Expect(Parse("codewars")).To(Equal([]int{0}))
	})
})
