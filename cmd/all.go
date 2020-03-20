package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	templateDir string
	outputDir   string
	overwrite   bool
	dryRun      bool
	format      bool
	examples    string
	gqlURL      string
	headers     []string

	allCmd = &cobra.Command{
		Use:   "all [OPTIONS]",
		Short: "A customizeable GraphQL markdown document generator",
		Long: `GraphQL Doc enables you to create customized markdown
		documentation driven by graphql query introspection on any valid graphql endpoint`,
	}
)

func init() {
	generateCmd.AddCommand(allCmd)

	allCmd.Flags().StringVarP(&gqlURL, "url", "u", "", "Location of customized markdown templates (default \"graphqldocs template\")")
	allCmd.Flags().StringArrayVarP(&headers, "header", "H", make([]string, 0), "Pass custom header to server e.g. \"Authorization: Bearer eyJh...\". This option can be used multiple times to add multiple headers.")
	allCmd.Flags().StringVarP(&examples, "examples", "x", "", "")
}
