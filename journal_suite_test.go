package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJournal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Journal Suite")
}
