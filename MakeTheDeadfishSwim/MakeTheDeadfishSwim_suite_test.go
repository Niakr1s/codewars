package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMakeTheDeadfishSwim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MakeTheDeadfishSwim Suite")
}
