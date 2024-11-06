package generator

import (
	"go/build"
	"log"
	"os"
	"path/filepath"

	"github.com/PlayerR9/mysd-lib/common"
)

// GoExt is the extension for Go files.
const GoExt string = ".go"

// GetPkgName returns the package name of the given file.
//
// Parameters:
//   - loc: The location of the file. This must be a Go file.
//
// Returns:
//   - string: The package name of the file.
//   - error: An error if getting the package name failed.
func GetPkgName(loc string) (string, error) {
	if loc == "" {
		return "", common.NewErrBadParam("loc", "must not be empty")
	}

	ext := filepath.Ext(loc)
	if ext != GoExt {
		return "", common.NewErrBadParam("loc", "must be a Go file")
	}

	dir_loc := filepath.Dir(loc)
	if dir_loc != "." {
		_, dir := filepath.Split(dir_loc)
		return dir, nil
	}

	pkg, err := build.ImportDir(loc, 0)
	if err != nil {
		return "", err
	}

	return pkg.Name, nil
}

// NewLogger creates a new logger with the given name.
//
// If the name is empty, it defaults to "generator".
//
// Parameters:
//   - name: The name of the logger.
//
// Returns:
//   - *log.Logger: The logger. Never returns nil.
func NewLogger(name string) *log.Logger {
	if name == "" {
		name = "generator"
	}

	return log.New(os.Stdout, "["+name+"]: ", log.Lshortfile)
}
