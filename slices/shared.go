package slices

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - slice: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices. Nil if no separator is found.
func IndicesOf[T comparable](slice []T, sep T) []int {
	if len(slice) == 0 {
		return nil
	}

	var count int

	for i := range slice {
		if slice[i] == sep {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	indices := make([]int, 0, count)

	for i := range slice {
		if slice[i] == sep {
			indices = append(indices, i)
		}
	}

	return indices
}
