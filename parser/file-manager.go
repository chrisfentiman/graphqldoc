package parser

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var (
	queryFile    = "query.md"
	objectFile   = "object.md"
	mutationFile = "mutation.md"
	scalarFile   = "scalar.md"
	enumFile     = "enum.md"
	ifaceFile    = "interface.md"
	inputFile    = "input.md"
)

// outFiles creates absolute paths to all of the markdown documents.
func outFiles(out string) *gqlFiles {

	dir, err := absolutePath(out)
	if err != nil {
		log.Fatalf("Unable to create an absolute path for out %s: %s", out, err)
	}

	return &gqlFiles{
		dir:      out,
		query:    filepath.Join(dir, queryFile),
		object:   filepath.Join(dir, objectFile),
		mutation: filepath.Join(dir, mutationFile),
		scalar:   filepath.Join(dir, scalarFile),
		enum:     filepath.Join(dir, enumFile),
		iface:    filepath.Join(dir, ifaceFile),
		input:    filepath.Join(dir, inputFile),
	}
}

// mkdir is a helper function for os.Mkdir
// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
// If there is an error, it will be of type *os.PathError.
func (d *docGenerator) mkdir() error {
	if _, err := os.Stat(d.outFiles.dir); !os.IsNotExist(err) && d.overwrite {
		os.RemoveAll(d.outFiles.dir)
		os.Remove(d.outFiles.dir)
	}

	return os.Mkdir(d.outFiles.dir, 0755)
}

// absolutePath is a helper function for filepath.Abs
// returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path.
func absolutePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		return filepath.Abs(path)
	}

	return path, nil
}

// relativePath is a helper function to turn any dir to a project level relative path
// Rel returns a relative path that is lexically equivalent to targpath when
// joined to basepath with an intervening separator.
func relativePath(dir string) string {
	base, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Unable to get home directory %s", err)
	}
	rel, err := filepath.Rel(base, dir)
	if err != nil {
		log.Fatalf("Unable to generate relative path %s", err)
	}

	return rel
}
