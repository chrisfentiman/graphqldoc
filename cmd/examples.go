package cmd

import (
	"github.com/spf13/cobra"
)

var (
	egCmd = &cobra.Command{
		Use:   "examples [OPTIONS]",
		Short: "A customizeable GraphQL markdown document generator",
		Long: `GraphQL Doc enables you to create customized markdown
		documentation driven by graphql query introspection on any valid graphql endpoint`,
	}
)

func init() {
	generateCmd.AddCommand(egCmd)

	egCmd.Flags().StringVarP(&examples, "examples", "x", "", "")
}
