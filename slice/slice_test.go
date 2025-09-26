package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	// Define a set of test cases with inputs and the expected output for each scenario.
	// The test cases cover a range of situations, including nil inputs, empty inputs, and
	// slices with elements, to ensure the Merge function behaves as intended in all cases.
	cases := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{name: "Both slices nil", first: nil, second: nil, expected: []int{}},
		{name: "First slice nil", first: nil, second: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "Second slice nil", first: []int{4, 5, 6}, second: nil, expected: []int{4, 5, 6}},
		{name: "Both slices empty", first: []int{}, second: []int{}, expected: []int{}},
		{name: "First slice empty", first: []int{}, second: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "Second slice empty", first: []int{4, 5, 6}, second: []int{}, expected: []int{4, 5, 6}},
		{name: "Both slices have elements", first: []int{7, 8, 9}, second: []int{10, 11, 12}, expected: []int{7, 8, 9, 10, 11, 12}},
		{name: "Slices with overlapping elements", first: []int{1, 2, 3}, second: []int{3, 4, 5}, expected: []int{1, 2, 3, 3, 4, 5}},
		{name: "Different length slices", first: []int{1}, second: []int{2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
	}

	// Iterate over the defined test cases, executing each one as a subtest.
	// Subtests allow each test case to be run independently, making it easier
	// to identify which specific case fails if an assertion does not hold.
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Call the Merge function with the input slices from the current test case.
			// The Merge function is expected to concatenate the two slices and return the result.
			result := Merge(tt.first, tt.second)

			// Compare the result of the Merge function with the expected output.
			// The assertion checks for equality and fails the test if the values do not match.
			assert.Equal(t, tt.expected, result, "Result for test case %q did not match expected output", tt.name)
		})
	}
}
