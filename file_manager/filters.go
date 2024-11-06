package file_manager

import "io/fs"

// RejectNilDirEntry filters out nil elements from a slice of fs.DirEntry pointers,
// modifying the slice in-place. If the slice is nil or empty, it does nothing.
//
// Parameters:
//   - entries: A pointer to a slice of fs.DirEntry pointers. The slice will be
//     modified to remove any nil elements. If all elements are nil, the slice
//     will be set to nil.
func RejectNilDirEntry(entries *[]fs.DirEntry) {
	if entries == nil || len(*entries) == 0 {
		return
	}

	var top int

	for _, entry := range *entries {
		if entry == nil {
			continue
		}

		(*entries)[top] = entry
		top++
	}

	if top == 0 {
		clear(*entries)
		*entries = nil
	} else if top == len(*entries) {
		return
	}

	clear((*entries)[top:])
	*entries = (*entries)[:top]
}

// RejectDir filters out nil elements and directory entries from a slice of fs.DirEntry
// pointers, modifying the slice in-place. If the slice is nil or empty, it does nothing.
//
// Parameters:
//   - entries: A pointer to a slice of fs.DirEntry pointers. The slice will be
//     modified to remove any nil or directory elements. If all elements are filtered out,
//     the slice will be set to nil.
func RejectDir(entries *[]fs.DirEntry) {
	if entries == nil || len(*entries) == 0 {
		return
	}

	var top int

	for _, entry := range *entries {
		if entry == nil || entry.IsDir() {
			continue
		}

		(*entries)[top] = entry
		top++
	}

	if top == 0 {
		clear(*entries)
		*entries = nil
	} else if top == len(*entries) {
		return
	}

	clear((*entries)[top:])
	*entries = (*entries)[:top]
}
