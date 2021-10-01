package lib

import (
	"github.com/jysandy/table-driven-tests/testhelpers"
	"testing"
)

func TestFoo(t *testing.T) {
	tests := []testhelpers.TableDrivenTest{
		{"Pineapples",
			[]interface{}{1, "pineapple"},
			"I see 1 pineapple."},
		{"Sheep",
			[]interface{}{69, "sheep"},
			"I see 69 sheep."},
		{"Whales",
			[]interface{}{42, "whales"},
			"I see 42 whales."},
	}

	testhelpers.RunTableDrivenTests(t, Foo, tests)
}
