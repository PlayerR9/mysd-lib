package strings

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - slice: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices. Nil if no separator is found.
func IndicesOf(slice []string, sep string) []int {
	if len(slice) == 0 {
		return nil
	}

	var count int

	for i := 0; i < len(slice); i++ {
		if slice[i] != sep {
			continue
		}

		count++
	}

	if count == 0 {
		return nil
	}

	indices := make([]int, 0, count)

	for i := 0; i < len(slice) && len(indices) < count; i++ {
		if slice[i] != sep {
			continue
		}

		indices = append(indices, i)
	}

	return indices
}
