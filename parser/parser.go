package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

var (
	private = "_"
)

// newGenerator initializes a new document generator configuration based on the
// values of the params.
func newGenerator(schema *Schema, templateDir string, format bool, overwrite bool, outDir string, dryRun bool) *docGenerator {

	if templateDir != "" {
		path, err := absolutePath(templateDir)
		if err != nil {
			log.Fatalf("Unable to create an absolute path for out %s: %s", templateDir, err)
		}

		templateDir = path
	}

	return &docGenerator{
		schema:    schema,
		templates: templateDir,
		format:    format,
		overwrite: overwrite,
		dryRun:    dryRun,
		outFiles:  outFiles(outDir),
	}
}

// generate is the main function - it generates
// the documentation based on the introspection query
// and the gqldoc files written to memory from either
// local or a user provided directory
func (d *docGenerator) generate() {

	if err := d.mkdir(); err != nil {
		log.Fatalf("Unable to create directory %s", err)
	}

	var (
		scalars []*FullType
		enums   []*FullType
		ifaces  []*FullType
		objects []*FullType
	)

	for _, v := range d.schema.Types {
		if !strings.Contains(v.Name, "_") {
			v.Fields = cleanType(v.Fields, private)
			switch v.Kind {
			case "SCALAR":
				scalars = append(scalars, v)
				break
			case "ENUM":
				enums = append(enums, v)
				break
			case "INTERFACE":
				ifaces = append(ifaces, v)
				break
			case "OBJECT":
				objects = append(objects, v)
				break
			}
		}
	}

	chanMerged := merge(d.fullType(d.schema.QueryType, query),
		d.fullType(d.schema.MutationType, mutation),
		d.fullTypes(scalars, scalar),
		d.fullTypes(enums, enum),
		d.fullTypes(ifaces, iface),
		d.fullTypes(objects, object))

	errs := make([]error, 0)
	for err := range chanMerged {
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, err := range errs {
		log.Printf("%s", err)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}

	os.Exit(0)

}

// merge merges all given channels into one channel
// creating a fan-in pattern.
func merge(outputsChan ...<-chan error) <-chan error {
	wg := &sync.WaitGroup{}

	merged := make(chan error, 100)

	wg.Add(len(outputsChan))

	output := func(wg *sync.WaitGroup, oc <-chan error) {
		for o := range oc {
			merged <- o
		}
		wg.Done()
	}

	for _, optChan := range outputsChan {
		go output(wg, optChan)
	}

	go func(wg *sync.WaitGroup, merged chan error) {
		wg.Wait()
		close(merged)
	}(wg, merged)

	return merged
}

// cleanType removes private types from
// graphql introspection query - these types
// will not be written to file.
func cleanType(tfs []*TypeField, comparable string) []*TypeField {
	fields := make([]*TypeField, 0)
	for _, field := range tfs {
		if !strings.Contains(field.Name, comparable) {
			fields = append(fields, field)
		}
	}

	return fields
}

// fullType writes *FullType documentation to file. FullTypes are generally
// Root level Query or Mutations.
func (d *docGenerator) fullType(ft *FullType, gqlt gqlType) <-chan error {
	ftChan := make(chan error, 100)

	go func(ft *FullType, gqlt gqlType, d *docGenerator, ftChan chan error) {
		defer close(ftChan)

		var (
			file string
		)

		switch gqlt {
		case query:
			file = d.outFiles.query
		case mutation:
			file = d.outFiles.mutation
		default:
			return
		}

		ft.Fields = cleanType(ft.Fields, private)

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ftChan <- fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
			return
		}

		gqldoc, err := getTemplate(d.templates, gqlt)
		if err != nil {
			ftChan <- fmt.Errorf("Unable to get %s template: %s", gqlt, err)
			return
		}

		t := template.Must(tempGen(d.outFiles.dir, gqldoc))
		err = t.Execute(f, ft)
		if err != nil {
			ftChan <- fmt.Errorf("TODO: %s %s", gqlt, err)
			return
		}
	}(ft, gqlt, d, ftChan)

	return ftChan
}

// fullType is like fullType except it writes []*FullType documentation to file. FullTypes are generally
// objects, scalars, input objects, enum, unions
func (d *docGenerator) fullTypes(fts []*FullType, gqlt gqlType) <-chan error {
	ftsChan := make(chan error, 100)

	go func(fts []*FullType, gqlt gqlType, d *docGenerator, ftsChan chan error) {
		defer close(ftsChan)

		var file string

		if len(fts) < 1 {
			return
		}

		switch gqlt {
		case scalar:
			file = d.outFiles.scalar
		case enum:
			file = d.outFiles.enum
		case object:
			file = d.outFiles.object

			modified := make([]*FullType, 0)
			for _, ft := range fts {
				switch ft.Name {
				case "Query", "Mutation":
				default:
					modified = append(modified, ft)
				}
			}

			fts = modified
		case iface:
			file = d.outFiles.iface
		// TODO: case input:
		// 	file = d.outFiles.input
		// 	gqls = "input"
		default:
			return
		}

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ftsChan <- fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		gqldoc, err := getTemplate(d.templates, gqlt)
		if err != nil {
			ftsChan <- fmt.Errorf("Unable to get %s template: %s", gqlt, err)
		}

		t := template.Must(tempGen(d.outFiles.dir, gqldoc))
		err = t.Execute(f, fts)
		if err != nil {
			ftsChan <- fmt.Errorf("TODO: %s %s", gqlt, err)
		}
	}(fts, gqlt, d, ftsChan)

	return ftsChan
}

// getTemplate retrieves .gqldoc files from a specified template dir
// or it will use the in memory templates.
func getTemplate(templateDir string, gqlt gqlType) (string, error) {
	switch gqlt {
	case query, mutation:
		if templateDir == "" {
			data, err := Asset("template/schema.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "schema.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case scalar:
		if templateDir == "" {
			data, err := Asset("template/scalar.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "scalar.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case enum:
		if templateDir == "" {
			data, err := Asset("template/enum.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "enum.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case object:
		if templateDir == "" {
			data, err := Asset("template/object.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "object.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case iface:
		if templateDir == "" {
			data, err := Asset("template/interface.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "interface.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case input:
		if templateDir == "" {
			data, err := Asset("template/input.gqldoc")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "input.gqldoc"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	default:
		return "", fmt.Errorf("Internal Error: raise issue in graphqldoc github, recieved: %s %v", templateDir, gqlt)
	}
}

// tempGen writes data to template in memory and returns back a
// written template.
//
// tempGen also creates template helper functions, such as transform
// which allows gqldoc writers to transform text into any of the following:
// lowercase, UPPERCASE, Title Case, Sentence case, PascalCase, camelCase, kebab-case, snake_case
//
//
func tempGen(dir string, data string) (*template.Template, error) {
	p, err := template.New("MD").Funcs(template.FuncMap{
		"transform": func(to string, str string) string {
			switch to {
			case "lower", "lowercase", "loc":
				return strings.ToLower(str)
			case "upper", "UPPERCASE", "upc":
				return strings.ToUpper(str)
			case "title", "Title Case", "tlc":
				return title(strings.ToLower(str))
			case "sentence", "Sentence case", "stc":
				return firstToUpper(strings.ToLower(str))
			case "pascal", "PascalCase", "psc":
				return runeMap(title(runeMap(str, []rune(" "), false)), []rune{}, true)
			case "camel", "camelCase", "cmc":
				// to lower and replace hyphens and underscores with spaces
				lowerd := runeMap(strings.ToLower(str), []rune(" "), false)
				return runeMap(firstToLower(title(lowerd)), []rune{}, true)
			case "kebab", "kebab-case", "kbc":
				return runeMap(strings.ToLower(str), []rune("-"), false)
			case "snake", "snake_case", "skc":
				return runeMap(strings.ToLower(str), []rune("_"), false)
			default:
				return str
			}
		},
		// TODO: revise this function
		"getType": func(t *TypeRef) interface{} {
			value := struct {
				Name string
				Type string
				Kind string
				Dir  string
			}{Type: "%s"}
			for t.OfType != nil {
				if t.Kind == "NON_NULL" {
					value.Type = value.Type + "!"
				}
				if t.Kind == "LIST" {
					value.Type = "[" + value.Type + "]"
				}
				t = t.OfType
			}
			value.Name = t.Name
			value.Kind = t.Kind
			value.Type = fmt.Sprintf(value.Type, value.Name)
			value.Name = strings.Replace(strings.ToLower(value.Name), " ", "-", -1)
			value.Dir = relativePath(dir)
			return value

		},
	}).Parse(data)
	if err != nil {
		log.Fatalf("TODO: %s", err)
	}
	return p, err
}
