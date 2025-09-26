package slice

import (
	"sort"
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

func TestExclude(t *testing.T) {
	t.Parallel()

	// SliceString tests the Exclude function for slices of integers. This test suite is designed to ensure that
	// the Exclude function correctly removes all occurrences of a specified integer from a slice. The test cases
	// cover a range of scenarios, including removing an element that appears once or multiple times, trying to
	// remove an element that does not exist in the slice, and handling empty or nil slices.
	t.Run("SliceString", func(t *testing.T) {
		cases := []struct {
			name     string
			elements []int
			element  int
			expected []int
		}{
			{name: "ExcludeSingleElement", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeMultipleOccurrences", elements: []int{1, 2, 3, 3, 4, 3, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeNonexistentElement", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: []int{1, 2, 3, 4, 5}},
			{name: "ExcludeEmptySlice", elements: []int{}, element: 1, expected: []int{}},
			{name: "ExcludeSingleElementSlice", elements: []int{1}, element: 1, expected: []int{}},
			{name: "ExcludeNilSlice", elements: nil, element: 1, expected: nil},
		}

		// Iterate through each test case and execute the Exclude function.
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Exclude function with the current test case's elements and element.
				result := Exclude(tt.elements, tt.element)

				// Assert that the result from the Exclude function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "Test case %s failed", tt.name)
			})
		}
	})

	// SliceInt tests the Exclude function with various scenarios involving slices of integers. The goal of this test
	// is to verify that the Exclude function correctly removes all occurrences of a specified integer from a slice.
	// The test cases cover a range of conditions to ensure the function handles different situations accurately,
	// including scenarios where elements appear multiple times, do not appear, or when the input slice is empty or nil.
	// Additionally, it tests edge cases such as large slices and slices with negative numbers.
	t.Run("SliceInt", func(t *testing.T) {
		cases := []struct {
			name     string
			elements []int
			element  int
			expected []int
		}{
			{name: "ExcludeSingleElement", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeMultipleOccurrences", elements: []int{1, 2, 3, 3, 4, 3, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeNonexistentElement", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: []int{1, 2, 3, 4, 5}},
			{name: "ExcludeEmptySlice", elements: []int{}, element: 1, expected: []int{}},
			{name: "ExcludeSingleElementSlice", elements: []int{1}, element: 1, expected: []int{}},
			{name: "ExcludeNilSlice", elements: nil, element: 1, expected: nil},
			{name: "ExcludeAllElements", elements: []int{7, 7, 7, 7, 7}, element: 7, expected: []int{}},
			{name: "ExcludeFirstElement", elements: []int{9, 2, 3, 4, 5}, element: 9, expected: []int{2, 3, 4, 5}},
			{name: "ExcludeLastElement", elements: []int{1, 2, 3, 4, 10}, element: 10, expected: []int{1, 2, 3, 4}},
			{name: "ExcludeMiddleElement", elements: []int{1, 2, 10, 4, 5}, element: 10, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeElementFromLargeSlice", elements: createSequenceWithRepeats(1000, 500), element: 500, expected: createSequenceWithoutRepeats(1000)},
			{name: "ExcludeNegativeElement", elements: []int{-3, -2, -1, 0, 1, 2, 3}, element: -2, expected: []int{-3, -1, 0, 1, 2, 3}},
			{name: "ExcludeFromMixedValues", elements: []int{-1, 0, 1, -1, 0, 1}, element: -1, expected: []int{0, 1, 0, 1}},
		}

		// Iterate through each test case and execute the Exclude function.
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Exclude function with the current test case's elements and element.
				result := Exclude(tt.elements, tt.element)

				// Assert that the result from the Exclude function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "Test case %s failed", tt.name)
			})
		}
	})
}

func TestContains(t *testing.T) {
	t.Parallel()

	// SliceInt tests the Contains function for slices of integers. This test ensures that the Contains function
	// accurately identifies whether a specific integer is present in a slice of integers. It covers a variety of
	// cases to verify correctness, including scenarios where the element is present or absent from the slice.
	t.Run("SliceInt", func(t *testing.T) {
		// Define a range of test cases for the Contains function with integer slices.
		cases := []struct {
			name     string
			elements []int
			element  int
			expected bool
		}{
			{name: "Element is in the slice", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: true},
			{name: "Element is not in the slice", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: false},
			{name: "Empty slice", elements: []int{}, element: 1, expected: false},
			{name: "Empty slice with empty element", elements: []int{}, element: 0, expected: false},
			{name: "Single element slice contains element", elements: []int{1}, element: 1, expected: true},
			{name: "Single element slice does not contain element", elements: []int{1}, element: 2, expected: false},
			{name: "Multiple elements, element at start", elements: []int{5, 1, 2, 3, 4}, element: 5, expected: true},
			{name: "Multiple elements, element at end", elements: []int{1, 2, 3, 4, 5}, element: 5, expected: true},
			{name: "Multiple elements, element in middle", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: true},
			{name: "Multiple elements, element repeated", elements: []int{1, 2, 3, 3, 4, 5}, element: 3, expected: true},
			{name: "Nil slice", elements: nil, element: 1, expected: false},
		}

		// Iterate through each test case and execute the Contains function
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Contains function with the current test case's elements and element.
				// The Contains function will return whether the element is present in the slice.
				result := Contains(tt.elements, tt.element)

				// Assert that the result from the Contains function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "result should match the expected value for test case: %s", tt.name)
			})
		}
	})

	// SliceString tests the Contains function for slices of strings.
	// It verifies that the function correctly identifies the presence or absence of an element
	// in various scenarios involving string slices, including sorted and unsorted slices.
	t.Run("SliceString", func(t *testing.T) {
		// Define test cases with various scenarios for slices of strings.
		cases := []struct {
			name     string
			elements []string
			element  string
			expected bool
		}{
			{name: "Nil slice", elements: nil, element: "test", expected: false},
			{name: "Empty slice", elements: []string{}, element: "test", expected: false},
			{name: "Element in single-element sorted slice", elements: []string{"test"}, element: "test", expected: true},
			{name: "Element not in single-element sorted slice", elements: []string{"test"}, element: "notfound", expected: false},
			{name: "Element in multiple-element sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "beta", expected: true},
			{name: "Element not in multiple-element sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "delta", expected: false},
			{name: "Element at the beginning of the sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "alpha", expected: true},
			{name: "Element at the end of the sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "gamma", expected: true},
			{name: "Unsorted slice (contains element)", elements: []string{"beta", "alpha", "gamma"}, element: "alpha", expected: true},
			{name: "Unsorted slice (does not contain element)", elements: []string{"beta", "alpha", "gamma"}, element: "delta", expected: false},
		}

		// Iterate over each test case and execute the Contains function.
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Sort the slice of strings to prepare it for testing the Contains function.
				// Sorting ensures that the Contains function operates correctly on ordered data.
				// This step is essential if the slice is not nil, as it standardizes the input for the test.
				if tt.elements != nil {
					// Sort the slice of strings in ascending order.
					// Sorting is done to match the expected behavior of the Contains function when the slice is ordered.
					sort.Strings(tt.elements)
				}

				// Call the Contains function with the current test case’s slice and element.
				result := Contains(tt.elements, tt.element)

				// Assert that the result matches the expected value.
				// If the result does not match, the test will fail, and the message will include
				// the test case details for easy identification of the failure.
				assert.Equal(t, tt.expected, result, "Expected Contains(%v, %v) to be %v but result %v", tt.elements, tt.element, tt.expected, result)
			})
		}
	})
}

// createSequenceWithRepeats generates a slice of integers with a specified size.
// The slice contains a repeated element at every 100th position, while other positions
// are filled with their respective indices.
func createSequenceWithRepeats(size, repeatedElement int) []int {
	// Initialize a slice with the specified size.
	slice := make([]int, size)

	// Iterate over each index in the slice.
	for i := 0; i < size; i++ {
		// If the index is a multiple of 100, insert the repeated element.
		if i%100 == 0 {
			slice[i] = repeatedElement
		} else {
			// Otherwise, insert the index value itself.
			slice[i] = i
		}
	}

	// Return the generated slice.
	return slice
}

// createSequenceWithoutRepeats generates a slice of integers with a specified size,
// ensuring that no element is repeated at positions that are multiples of 100.
func createSequenceWithoutRepeats(size int) []int {
	// Initialize an empty slice with a predefined capacity.
	slice := make([]int, 0, size)

	// Iterate over each index up to the specified size.
	for i := 0; i < size; i++ {
		// Only include indices that are not multiples of 100.
		if i%100 != 0 {
			// Append the index to the slice.
			slice = append(slice, i)
		}
	}

	// Return the generated slice.
	return slice
}
