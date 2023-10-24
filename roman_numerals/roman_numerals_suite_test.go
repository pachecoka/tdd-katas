package roman_numerals_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRomanNumerals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RomanNumerals Suite")
}
