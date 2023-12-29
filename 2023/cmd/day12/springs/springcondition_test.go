package springs_test

import (
	"fmt"
	"testing"
)

func TestGetRegexpString(t *testing.T) {
	cases := struct {
		input any
		want  any
	}{
		input: "foo",
		want:  "bar",
	}
	fmt.Println(cases)
}

// NO GOOD. Apparently, the XXXX ExampleXXXX must match something in the
// package being tested. Hmmm...
// func ExampleRegexpString() {
// 	fmt.Println(springs.SC_OPERATIONAL.GetRegexpString())
// 	// Output: [.?]
// }
