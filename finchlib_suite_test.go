package finchlib_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFinchlib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Finchlib Suite")
}
