package gotesting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGotesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotesting Suite")
}
