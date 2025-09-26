package slice

// Merge concatenates two slices into a single slice.
// It creates a new slice with a length equal to the sum of the lengths of the input slices.
// The function copies all elements from the first slice followed by all elements from the second slice into the new slice,
// and returns this combined slice.
func Merge[T any](first, second []T) []T {
	// Allocate a new slice with enough capacity to hold all elements from both input slices.
	list := make([]T, len(first)+len(second))
	// Copy all elements from the first slice into the new slice.
	copy(list, first)
	// Copy all elements from the second slice into the new slice, starting right after the first slice's elements.
	copy(list[len(first):], second)
	// Return the combined slice containing elements from both input slices.
	return list
}

// Exclude removes all instances of a specified value from the provided slice.
// It creates a new slice containing only the elements that are not equal to the specified value.
// This approach efficiently constructs the result slice by reusing the original slice's underlying array,
// avoiding unnecessary memory allocations.
func Exclude[T comparable](elements []T, element T) []T {
	// Initialize the result slice with the same underlying array as the original slice.
	// This avoids unnecessary allocations and keeps the capacity the same.
	result := elements[:0]

	// Iterate over each item in the original slice.
	for _, item := range elements {
		// Check if the current item is not equal to the specified value to be excluded.
		if item != element {
			// Append the item to the result slice if it is not equal to the specified value.
			result = append(result, item)
		}
	}

	// Return the filtered slice with the specified value removed.
	return result
}
