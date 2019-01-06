package gocounting

// Source : exercism/problem-specifications
// Commit :
// Problem Specifications Version: 1.0.0

var testCases = []struct {
	description    string
	board          string
	input          [2]int
	expectedPlayer string
	expectedArea   [][]int
}{
	{
		description:    "Test",
		board:          "",
		input:          [2]int{0, 0},
		expectedPlayer: "BLACK",
		expectedArea: [][]int{
			{1, 1}, {2, 2},
		},
	},
}
