package bytes

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - slice: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices. Nil if no separator is found.
func IndicesOf(slice []byte, sep []byte) []int {
	lenSep := len(sep)
	if lenSep == 0 {
		return nil
	}

	lenSlice := len(slice)
	if lenSlice == 0 {
		return nil
	}

	var count int

	for i := range slice {
		if slice[i] == sep[0] {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	indices := make([]int, 0, count)

	for i := range slice[:lenSlice-lenSep+1] {
		if slice[i] == sep[0] {
			indices = append(indices, i)
		}
	}

	var top int

	for i := 1; i < lenSep; i++ {
		top = 0

		for _, idx := range indices {
			if slice[idx+1] == sep[i] {
				indices[top] = idx
				top++
			}
		}

		if top == 0 {
			return nil
		}

		indices = indices[:top]
	}

	indices = indices[:len(indices):len(indices)]

	return indices
}
