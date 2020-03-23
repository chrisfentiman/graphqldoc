package parser

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
	input    string
	mutation string
	scalar   string
	enum     string
	iface    string
}

type Data struct {
	Schema *Schema `json:"__schema"`
}

// Response estructura de una respuesta HTTP
type Response struct {
	Data   *Data                    `json:"data"`
	Errors []map[string]interface{} `json:"errors"`
}

type IntrospectionQuery struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
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
