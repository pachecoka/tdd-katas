package bank_account_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBankAccount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankAccount Suite")
}
