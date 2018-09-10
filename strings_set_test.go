package finchlib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sparkymat/finchlib/strings"
)

var _ = Describe("strings.Set", func() {
	var (
		set *strings.Set
	)

	BeforeEach(func() {
		set = strings.NewSet()
	})

	It("creates a valid set", func() {
		Expect(set).ToNot(BeNil())
	})

	It("successfully adds to the set", func() {
		Expect(set.Includes("foo")).To(BeFalse())
		set.Add("foo")
		Expect(set.Includes("foo")).To(BeTrue())
	})

	It("successfully removes from the set", func() {
		set.Add("foo")
		Expect(set.Includes("foo")).To(BeTrue())
		set.Remove("foo")
		Expect(set.Includes("foo")).To(BeFalse())
	})

	It("returs a valid list", func() {
		set.Add("foo", "bar", "foo")
		Expect(set.List().Sort().Slice()).To(Equal([]string{"bar", "foo"}))
	})
})
