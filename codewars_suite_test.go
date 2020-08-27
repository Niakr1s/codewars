package codewars_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodewars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Codewars Suite")
}
