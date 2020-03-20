package cmd

import (
	"fmt"
	"net/url"

	"github.com/chrisfentiman/graphqldoc/parser"
	"github.com/spf13/cobra"
)

var (
	docsCmd = &cobra.Command{
		Use:   "docs [OPTIONS]",
		Short: "A customizeable GraphQL markdown document generator",
		Long: `GraphQL Doc enables you to create customized markdown
		documentation driven by graphql query introspection on any valid graphql endpoint`,
		Run: func(cmd *cobra.Command, args []string) {
			docsRunner(gqlURL, headers, templateDir, format, overwrite, outputDir, dryRun)
		},
	}
)

func init() {
	generateCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringVarP(&gqlURL, "url", "u", "", "Location of customized markdown templates (default \"GraphQL Doc Templates\")")
	docsCmd.Flags().StringArrayVarP(&headers, "header", "H", make([]string, 0), "Pass custom header to server e.g. \"Authorization: Bearer eyJh...\". This option can be used multiple times to add multiple headers.")
}

func docsRunner(endpoint string, headers []string, templates string, format bool, overwrite bool, out string, dryRun bool) {
	u, err := url.ParseRequestURI(endpoint)
	if err == nil {
		fmt.Print("Generating documentation...\n")
		parser.HTTP(u.String(), headers, templates, format, overwrite, out, dryRun)
		return
	}
}
