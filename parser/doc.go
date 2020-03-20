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

type TypeRef struct {
	Kind   string   `json:"kind"`
	Name   string   `json:"name"`
	OfType *TypeRef `json:"ofType"`
}

type InputValue struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	DefaultValue interface{} `json:"defaultValue"`
	Type         *TypeRef    `json:"type"`
}

type TypeField struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	Args              []*InputValue `json:"args"`
	Type              *TypeRef      `json:"type"`
	IsDeprecated      bool          `json:"isDeprecated"`
	DeprecationReason string        `json:"deprecationReason"`
}

type EnumValues struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

type FullType struct {
	Kind          string        `json:"kind"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Fields        []*TypeField  `json:"fields"`
	InputFields   []*InputValue `json:"inputFields"`
	Interfaces    []*TypeRef    `json:"interfaces"`
	EnumValues    []*EnumValues `json:"enumValues"`
	PossibleTypes []*TypeRef    `json:"possibleTypes"`
}

type TypeDirective struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Args        []*InputValue `json:"args"`
	OnOperation bool          `json:"onOperation"`
	onFragment  bool          `json:"onFragment"`
	onField     bool          `json:"onField"`
}

type Schema struct {
	QueryType        *FullType        `json:"queryType"`
	MutationType     *FullType        `json:"mutationType"`
	SubscriptionType *FullType        `json:"subscriptionType"`
	Types            []*FullType      `json:"types"`
	Directives       []*TypeDirective `json:"directives"`
}

type docGenerator struct {
	schema    *Schema
	templates string
	format    bool
	overwrite bool
	dryRun    bool
	outFiles  *gqlFiles
}

type gqlFiles struct {
	dir      string
	query    string
	object   string
	mutation string
	scalar   string
	enum     string
	iface    string
}

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

func (d *docGenerator) genDir() {
	if _, err := os.Stat(d.outFiles.dir); !os.IsNotExist(err) && d.overwrite {
		os.RemoveAll(d.outFiles.dir)
		os.Remove(d.outFiles.dir)
	}

	os.Mkdir(d.outFiles.dir, 0755)
}

func (d *docGenerator) generateDocs() {

	d.genDir()
	var (
		scalars []*FullType
		enums   []*FullType
		ifaces  []*FullType
		objects []*FullType
	)

	for _, v := range d.schema.Types {
		if !strings.Contains(v.Name, "_") {
			v.Fields = cleanType(v.Fields)
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

func cleanType(tfs []*TypeField) []*TypeField {
	fields := make([]*TypeField, 0)
	for _, field := range tfs {
		if !strings.Contains(field.Name, "_") {
			fields = append(fields, field)
		}
	}

	return fields
}

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

		ft.Fields = cleanType(ft.Fields)

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ftChan <- fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
			return
		}

		tmpl, err := getTemplate(d.templates, gqlt)
		if err != nil {
			ftChan <- fmt.Errorf("Unable to get %s template: %s", gqlt, err)
			return
		}

		t := template.Must(tempGen(d.outFiles.dir, tmpl))
		err = t.Execute(f, ft)
		if err != nil {
			ftChan <- fmt.Errorf("TODO: %s %s", gqlt, err)
			return
		}
	}(ft, gqlt, d, ftChan)

	return ftChan
}

func (d *docGenerator) fullTypes(fts []*FullType, gqlt gqlType) <-chan error {
	ftsChan := make(chan error, 100)

	go func(fts []*FullType, gqlt gqlType, d *docGenerator, ftsChan chan error) {
		defer close(ftsChan)

		var (
			file string
		)

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

		tmpl, err := getTemplate(d.templates, gqlt)
		if err != nil {
			ftsChan <- fmt.Errorf("Unable to get %s template: %s", gqlt, err)
		}

		t := template.Must(tempGen(d.outFiles.dir, tmpl))
		err = t.Execute(f, fts)
		if err != nil {
			ftsChan <- fmt.Errorf("TODO: %s %s", gqlt, err)
		}
	}(fts, gqlt, d, ftsChan)

	return ftsChan
}

type gqlType string

const (
	query    gqlType = "query"
	mutation gqlType = "mutation"
	scalar   gqlType = "scalar"
	enum     gqlType = "enum"
	object   gqlType = "object"
	iface    gqlType = "interface"
	input    gqlType = "input"
)

func getTemplate(templateDir string, gqlt gqlType) (string, error) {
	switch gqlt {
	case query, mutation:
		if templateDir == "" {
			data, err := Asset("template/schema.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "schema.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case scalar:
		if templateDir == "" {
			data, err := Asset("template/scalar.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "scalar.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case enum:
		if templateDir == "" {
			data, err := Asset("template/enum.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "enum.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case object:
		if templateDir == "" {
			data, err := Asset("template/object.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "object.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case iface:
		if templateDir == "" {
			data, err := Asset("template/interface.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "interface.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	case input:
		if templateDir == "" {
			data, err := Asset("template/input.tmpl")
			if err != nil {
				return "", err
			}

			return string(data), nil
		}

		dat, err := ioutil.ReadFile(filepath.Join(templateDir, "input.tmpl"))
		if err != nil {
			return "", fmt.Errorf("Unable to open %s markdown file: %s", gqlt, err)
		}

		return string(dat), nil
	default:
		return "", fmt.Errorf("Internal Error: raise issue in graphqldoc github, recieved: %s %v", templateDir, gqlt)
	}
}

func tempGen(dir string, data string) (*template.Template, error) {
	p, err := template.New("MD").Funcs(template.FuncMap{
		"getType": func(t *TypeRef) interface{} {
			value := struct {
				Name string
				Type string
				Kind string
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
			if t.Kind == "SCALAR" {
				value.Name = dir + "scalar#" + value.Name
			}
			if t.Kind == "OBJECT" {
				value.Name = dir + "object#" + value.Name
			}
			value.Name = strings.Replace(strings.ToLower(value.Name), " ", "-", -1)
			return value

		},
	}).Parse(data)
	if err != nil {
		log.Fatalf("TODO: %s", err)
	}
	return p, err
}
