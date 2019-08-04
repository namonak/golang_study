package bdd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getMax 함수 검증", func() {
	Context("Max", func() {
		It("주어진 정수 가운데 가장 큰 수를 반환한다.", func() {
			Expect(getMaxInt32(777, 100, 30)).Should(BeNumerically("==", 777))
			Expect(getMaxInt32(32, 29, 300)).Should(BeNumerically("==", 300))
			Expect(getMaxInt32(5, 7, 2)).Should(BeNumerically("==", 7))
		})
	})
})
