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
)

func outFiles(out string) *gqlFiles {

	dir := getAbs(out, false)
	return &gqlFiles{
		dir:      out,
		query:    filepath.Join(dir, queryFile),
		object:   filepath.Join(dir, objectFile),
		mutation: filepath.Join(dir, mutationFile),
		scalar:   filepath.Join(dir, scalarFile),
		enum:     filepath.Join(dir, enumFile),
		iface:    filepath.Join(dir, ifaceFile),
	}
}

func getAbs(path string, ignoreEmpty bool) string {
	if ignoreEmpty && path == "" {
		return path
	}

	dir := path
	if !filepath.IsAbs(dir) {
		abs, err := filepath.Abs(dir)
		if err != nil {
			log.Fatalf("Unable to create an absolute path for out %s: %s", path, err)
		}

		dir = abs
	}

	return dir
}

func (d *docGenerator) mkdir() error {
	if _, err := os.Stat(d.outFiles.dir); !os.IsNotExist(err) && d.overwrite {
		os.RemoveAll(d.outFiles.dir)
		os.Remove(d.outFiles.dir)
	}

	return os.Mkdir(d.outFiles.dir, 0755)
}

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
