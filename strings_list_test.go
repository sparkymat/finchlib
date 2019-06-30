package finchlib_test

import (
	gostrings "strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sparkymat/finchlib/strings"
)

var _ = Describe("StringsList", func() {
	var (
		list *strings.List
	)

	BeforeEach(func() {
		list = strings.NewList()
	})

	It("creates a valid empty list", func() {
		Expect(list).ToNot(BeNil())
	})

	It("returns a slice with the contained elements", func() {
		list.Add("foo", "bar")
		Expect(list.Slice()).To(Equal([]string{"foo", "bar"}))
		Expect(list.Slice()).ToNot(Equal([]string{"bar", "foo"}))
	})

	It("creates a valid list from a slice", func() {
		list = strings.ListFromSlice([]string{"foo", "bar"})
		Expect(list).ToNot(BeNil())
		Expect(list.Slice()).To(Equal([]string{"foo", "bar"}))
	})

	Describe("ListFromCSV", func() {
		It("creates a valid list from a csv without trimming space", func() {
			list = strings.ListFromCSV(" foo ,bar ", false)
			Expect(list).ToNot(BeNil())
			Expect(list.Slice()).To(Equal([]string{" foo ", "bar "}))
		})

		It("creates a valid list from a csv with trimming of space", func() {
			list = strings.ListFromCSV(" foo ,bar ", true)
			Expect(list).ToNot(BeNil())
			Expect(list.Slice()).To(Equal([]string{"foo", "bar"}))
		})
	})

	Describe("ListFromPgStringArrayValue", func() {
		It("returns error if the param is not a valid pg string array value", func() {
			list, err := strings.ListFromPgStringArray("something,else")
			Expect(list).To(BeNil())
			Expect(err).To(HaveOccurred())
		})

		It("returns a valid list if the param is valid", func() {
			list, err := strings.ListFromPgStringArray("{foo,bar}")
			Expect(list).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Slice()).To(Equal([]string{"foo", "bar"}))
		})
	})

	It("checks if element exists in the list", func() {
		Expect(list.Includes("foo")).To(BeFalse())
		list.Add("foo")
		Expect(list.Includes("foo")).To(BeTrue())
	})

	It("returns length of list", func() {
		list = strings.ListFromCSV("foo,bar", true)
		Expect(list.Length()).To(Equal(2))
	})

	It("adds element to the list", func() {
		list.Add("foo")
		Expect(list.Slice()).To(Equal([]string{"foo"}))
	})

	It("removes all occurrences in the list", func() {
		list = strings.ListFromCSV("foo,foo", true)
		Expect(list.Slice()).To(Equal([]string{"foo", "foo"}))

		list.RemoveAll("foo")
		Expect(list.Slice()).To(Equal([]string{}))
	})

	It("selects elements matching function", func() {
		list = strings.ListFromCSV("adam, betty, charlie, dory", true)
		newList := list.Select(func(e string) bool {
			if len(e) == 4 {
				return true
			}

			return false
		})
		Expect(newList.Slice()).To(Equal([]string{"adam", "dory"}))
	})

	It("rejects elements matching function", func() {
		list = strings.ListFromCSV("adam, betty, charlie, dory", true)
		newList := list.Reject(func(e string) bool {
			if e[0:1] == "a" || e[0:1] == "b" {
				return true
			}

			return false
		})
		Expect(newList.Slice()).To(Equal([]string{"charlie", "dory"}))
	})

	It("deletes based on index", func() {
		list = strings.ListFromCSV("foo, bar, jam", true)
		Expect(list.Slice()).To(Equal([]string{"foo", "bar", "jam"}))
		list.Delete(1)
		Expect(list.Slice()).To(Equal([]string{"foo", "jam"}))
	})

	It("iterates over the list", func() {
		list = strings.ListFromCSV("foo,bar,jam", true)

		iteratedList := []string{}
		list.Each(func(element string) {
			iteratedList = append(iteratedList, element)
		})

		Expect(iteratedList).To(Equal([]string{"foo", "bar", "jam"}))
	})

	It("maps the list using the provided function", func() {
		list = strings.ListFromCSV("foo,bar,jam", true)
		mappedList := list.Map(func(element string) string {
			return gostrings.ToUpper(element)
		})
		Expect(mappedList.Slice()).To(Equal([]string{"FOO", "BAR", "JAM"}))
	})

	It("returns csv for the list", func() {
		list = strings.ListFromCSV(" foo, bar,jam ", true)
		Expect(list.CSV()).To(Equal("foo,bar,jam"))
	})

	It("returns a joined string for the list", func() {
		list = strings.ListFromSlice([]string{"foo", "bar", "jam"})
		Expect(list.Join(":")).To(Equal("foo:bar:jam"))
	})

	It("returns a sorted list", func() {
		list = strings.ListFromCSV("foo,bar,jam", true)
		Expect(list.Sort().Slice()).To(Equal([]string{"bar", "foo", "jam"}))
	})

	It("retuns set for the set", func() {
		list = strings.ListFromCSV(" foo, foo, foo, bar", true)
		set := list.Set()
		Expect(set.List().Sort().Slice()).To(Equal([]string{"bar", "foo"}))
	})

	It("retusn pg_string_array value", func() {
		list = strings.ListFromSlice([]string{"foo", "bar"})
		Expect(list.PgStringArrayValue()).To(Equal("{foo,bar}"))
	})
})
