package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pachecoka/tdd-katas/fizzbuzz"
	"strconv"
)

var foo fizzbuzz.FizzbuzzUsecaseI = &fizzbuzz.FizzbuzzUsecase{}
var _ = Describe("Fizzbuzz", func() {
	When("Number is multiple of 3", func() {
		number := 3
		It("Should return Fizz", func() {
			Expect(foo.ExecuteFizzbuzzUsecase(number)).To(ContainElement("Fizz"))
		})
	})
	When("Number is multiple of 5", func() {
		number := 5
		It("Should return Buzz", func() {
			Expect(foo.ExecuteFizzbuzzUsecase(number)).To(ContainElement("Buzz"))
		})
	})
	When("Number is multiple of 7", func() {
		number := 7
		It("Should return Whizz", func() {
			Expect(foo.ExecuteFizzbuzzUsecase(number)).To(ContainElement("Whizz"))
		})
	})
	When("Number is multiple of 11", func() {
		number := 11
		It("Should return Bang", func() {
			Expect(foo.ExecuteFizzbuzzUsecase(number)).To(ContainElement("Bang"))
		})
	})
	When("Number is not multiple of 3,5,7,11", func() {
		number := 1
		It("Should return Fizz", func() {
			Expect(foo.ExecuteFizzbuzzUsecase(number)).To(Equal(strconv.Itoa(number)))
		})
	})
})
