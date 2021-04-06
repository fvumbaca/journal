package fs

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FS naming rules", func() {

	Context("on Jan 2", func() {
		baseDirectory := "/tmp/journal"
		var now time.Time

		BeforeEach(func() {
			now, _ = time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
		})

		It("should open the correct file for a daily note", func() {
			Expect(DayFilename(baseDirectory, now)).To(Equal("/tmp/journal/2006/January/02.md"))
		})
	})

})
