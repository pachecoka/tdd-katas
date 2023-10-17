package strange_characters_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStrangeCharacters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StrangeCharacters Suite")
}
