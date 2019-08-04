package bdd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_getMax(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integer Suite")
}
