package file_manager

import (
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/PlayerR9/mysd-lib/common"
)

// Exists returns whether the file at path exists and an error if any.
//
// Parameters:
//   - path: The path to the file.
//
// Returns:
//   - bool: True if the file exists, false if it does not.
//   - error: An error if the operation failed.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

// ExistsSpecific returns whether the file at path exists and an error if any.
// If want_dir is true, returns true only if the file is a directory.
// If want_dir is false, returns true only if the file is not a directory.
//
// Parameters:
//   - loc: The location of the file.
//   - want_dir: Whether the file must be a directory.
//
// Returns:
//   - bool: True if the file exists and matches the want_dir requirement, false if it does not.
//   - error: An error if the operation failed.
func ExistsSpecific(loc string, want_dir bool) (bool, error) {
	stat, err := os.Stat(loc)
	if err == nil {
		return want_dir && stat.IsDir(), nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return false, nil
}

// ScanDir traverses the directory located at 'loc' and populates 'files' with paths
// of files that satisfy the predicate 'p'. The traversal is performed iteratively
// using a stack to avoid stack overflow issues with deep directory structures.
//
// Parameters:
//   - loc: The root directory location where the scan begins.
//   - files: A pointer to a slice that will be populated with file paths that satisfy 'p'.
//   - p: A predicate function that determines if a file path should be included in 'files'.
//
// Returns:
//   - error: Returns an error if 'files' is nil or if there is an error reading a directory.
//
// Deprecated: This is not used anywhere in the codebase.
func ScanDir(loc string, files *[]string, p func(path string) bool) error {
	if p == nil {
		return nil
	} else if files == nil {
		return common.NewErrNilParam("files")
	}

	stack := []string{loc}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		entries, err := os.ReadDir(loc)
		if err != nil {
			return err
		}

		var dirs []string

		for _, entry := range entries {
			path := filepath.Join(top, entry.Name())

			if entry.IsDir() {
				_ = MayInsert(&dirs, path)
			} else if p(path) {
				_ = MayInsert(files, path)
			}
		}

		if len(dirs) > 0 {
			slices.Reverse(dirs)
			stack = append(stack, dirs...)
		}
	}

	return nil
}
