package parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

// HTTP execute query to the GraphQL endpoint
func HTTP(endpoint string, headers []string, templates string, format bool, overwrite bool, out string, dryRun bool) {
	var response Response

	query, err := Asset("template/schema.graphql")
	if err != nil {
		log.Fatalf("Couldn't get introspection query from memory: %s", err)
	}

	client := &http.Client{}
	b, err := json.Marshal(&IntrospectionQuery{Query: string(query)})
	if err != nil {
		log.Fatalf("Couldn't generate introspection query: %s", err)
	}

	// POST Request - As this is more reliable of an implementation than a GET due to the nature of GRAPHQL
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("Unable to initialize new HTTP request: %s", err)
	}

	req.Header.Add("content-type", "application/json")
	if len(headers) > 0 {
		for _, header := range headers {
			i := strings.Index(header, ":")
			name := header[:i]
			value := header[i+1 : len(header)]

			// Need to remove new lines created by some scripts due to long access tokens and trim all leading and trailing white space
			req.Header.Add(strings.TrimSpace(name), strings.ReplaceAll(strings.TrimSpace(value), "\n", ""))
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to complete POST request: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body from HTTP: %s", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Unable to parse introspection query: %s response: %s", err, string(body))
	}

	if err := resp.Body.Close(); err != nil {
		log.Fatalf("Error closing response object: %s", err)
	}

	if response.Data == nil && response.Errors == nil {
		log.Fatalf("Unexpected response from GraphQL: %s", string(body))
	}

	if len(response.Errors) > 0 {
		log.Fatalf("Errors recieved when completing introspection query: %v", response.Errors)
	}

	docs := &docGenerator{
		schema:    response.Data.Schema,
		templates: getAbs(templates, true),
		format:    format,
		overwrite: overwrite,
		dryRun:    dryRun,
		outFiles:  outFiles(out),
	}

	docs.generateDocs()
}
