package retcode_test

import (
	"fmt"
	"testing"

	"github.com/cardil/k1s/pkg/cli/retcode"
)

func TestCalc(t *testing.T) {
	cases := testCases()
	for i := range cases {
		tt := cases[i]
		t.Run(tt.name, func(t *testing.T) {
			if got := retcode.Calc(tt.err); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCases() []testCase {
	return []testCase{{
		name: "nil",
		err:  nil,
		want: 0,
	}, {
		name: "errExample",
		err:  errExample,
		want: 133,
	}, {
		name: "error of wrap caused by 12345",
		err:  fmt.Errorf("%w: 12345", errExample),
		want: 249,
	}}
}

var errExample = fmt.Errorf("example error")

type testCase struct {
	name string
	err  error
	want int
}
