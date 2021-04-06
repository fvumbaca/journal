package fs

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FS Suite")
}
