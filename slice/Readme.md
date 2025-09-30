# Slice Utility Package

This Go package provides a set of generic utility functions for working with slices. It leverages Go's generics to offer type-safe operations for common slice manipulations, such as merging, filtering, mapping, and removing duplicates. The package is designed to be simple, efficient, and reusable across various types.

## Installation

To use this package in your Go project, ensure you have Go 1.18 or later (required for generics). You can include the package by importing it:

```go
import (
    "github.com/spacemagneto/common/slice"
)
```

If this package is hosted in a repository, you can add it using:

```bash
  go get github.com/spacemagneto/common/slice
```

## Features

> The package includes the following functions:


- **Merge[T any](first, second []T) []T**: Concatenates two slices into a single slice.

- **Exclude[T comparable](elements []T, element T) []T**: Removes all instances of a specified value from a slice.

- **Contains[T constraints.Ordered](elements []T, element T) bool**: Checks if a slice contains a specific element using binary search.

- **Map[A, B any](elements []A, fn func(A) B) []B**: Applies a transformation function to each element of a slice, returning a new slice with transformed values.

- **Filter[T any](elements []T, fn func(T) bool) []T**: Filters a slice based on a predicate function, returning a new slice with elements that satisfy the predicate.

- **Unique[T comparable](elements []T) []T**: Removes duplicate elements from a slice, preserving the original order.


## Usage Examples

Below are examples demonstrating how to use each function in the package.

> ### Merge

Combines two slices into a single slice.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    a := []int{1, 2, 3}
    b := []int{4, 5, 6}
    result := slice.Merge(a, b)
    fmt.Println(result) // Output: [1 2 3 4 5 6]
}
```

> ### Exclude

Removes all instances of a specified value from a slice.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    numbers := []int{1, 2, 2, 3, 2, 4}
    result := slice.Exclude(numbers, 2)
    fmt.Println(result) // Output: [1 3 4]
}
```

> ### Contains

Checks if a slice contains a specific element using binary search.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    numbers := []int{3, 1, 4, 1, 5}
    found := slice.Contains(numbers, 4)
    fmt.Println(found) // Output: true

    found = slice.Contains(numbers, 6)
    fmt.Println(found) // Output: false
}
```

> ### Map

Transforms each element in a slice using a provided function.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    numbers := []int{1, 2, 3}
    doubled := slice.Map(numbers, func(n int) int { return n * 2 })
    fmt.Println(doubled) // Output: [2 4 6]
}
```

> ### Filter

Filters a slice to include only elements that satisfy a predicate function.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    evens := slice.Filter(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens) // Output: [2 4]
}
```

> ### Unique

Removes duplicate elements from a slice, preserving the original order.

```go
package main

import (
    "fmt"
    "github.com/spacemagneto/common/slice"
)

func main() {
    numbers := []int{1, 2, 2, 3, 3, 4}
    unique := slice.Unique(numbers)
    fmt.Println(unique) // Output: [1 2 3 4]
}
```

> ## Notes

- The package is designed to work with Go 1.18+ due to its use of generics.
- The Contains function sorts a copy of the input slice to perform a binary search, which may not be ideal for very large slices. Consider pre-sorting the slice if you need to call Contains multiple times.
- All functions return new slices to avoid modifying the input slices, ensuring immutability and thread safety. 
- The package is lightweight and has no external runtime dependencies beyond the Go standard library and golang.org/x/exp/constraints.

# License

This package is licensed under the Apache License, Version 2.0. See the LICENSE file for details.
