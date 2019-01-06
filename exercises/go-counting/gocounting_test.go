package gocounting

import (
	"testing"
)

// A stub file describing the desired API is given

func TestCount(t *testing.T) {
	for _, c := range testCases {
		if p, a := Board(c.board).Count(input[0], input[1]); p != c.expectedPlayer && a != c.expectedArea {
			t.Fatalf("Fail !")
		}

		t.Logf("Pass !")
	}
}
