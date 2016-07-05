package gotesting_test

import (
	. "github.com/ebonet/gotesting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/extensions/table"
	"fmt"
)

var _ = Describe("Blackjack", func() {

	Describe("Scoring Hand", func() {
		Context("with empty hand", func() {
			It("Should return 0", func() {
				Expect(ComputeScore("")).To(Equal(0))
			})
		})

		Context("with hand A2", func() {
			It("Should return 13", func() {
				Expect(ComputeScore("A2")).To(Equal(13))
			})
		})
	})

	DescribeTable("Scoring a hand",
		func(hand string, expected int) {
			Expect(ComputeScore(hand)).To(Equal(expected))
		},
		Entry("2","2", 2),
		Entry("3","3", 3),
		Entry("J","J", 10),
		Entry("A","A", 11),
		Entry("AJ","AJ", 21),
		Entry("A345","A345", 23),
		Entry("Invalid Char","/345≈", 0), // Invalid character
		Entry("Empty Hand", "", 0),      // Empty hand
	)

	Describe("Testing for blackjack", func() {
		contexts := []struct {
			Hand           string
			ExpectedResult bool
		}{
			{"AJ", true},
			{"JA", true},
			{"A4", false},
			{"J2", false},
			{"2", false},
			{"AJ5", false},
		}

		for _, ctx := range contexts {
			ctx := ctx //gotcha!
			Context("With hand"+ctx.Hand, func(){
				It(fmt.Sprintf("Should return %b", ctx.ExpectedResult), func() {
					Expect(CheckIfBlackjack(ctx.Hand)).To(Equal(ctx.ExpectedResult))
				})
			})
		}
	})
})
