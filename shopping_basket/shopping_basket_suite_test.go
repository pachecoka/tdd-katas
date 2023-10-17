package shopping_basket_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShoppingBasket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShoppingBasket Suite")
}
