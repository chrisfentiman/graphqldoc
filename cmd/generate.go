package cmd

import (
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate [COMMAND] [OPTIONS]",
		Short: "Runs the graphql document generator",
	}
)

// graphqldoc generate all -gql http://localhost:3000/graphql -t "./docs/generator/graphql" -eg "./src/**/*.ts"
// graphqldoc generate docs -gql http://localhost:3000/graphql -f "./docs/generator/graphql"
// graphqldoc generate examples -eg "./src/**/*.ts" -o "./docs/graphql"

// @graphqldocs:start
// type Example {
// 	string: String!
// }
// @graphqldocs:end
// Execute executes the root command.

func init() {
	rootCmd.AddCommand(generateCmd)
}
