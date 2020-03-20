package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "graphqldoc [COMMANDS] [OPTIONS]",
		Short: "A customizeable GraphQL markdown document generator",
		Long: `GraphQL Doc enables you to create customized markdown
		documentation driven by graphql query introspection on any valid graphql endpoint`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&templateDir, "template-dir", "t", "", "Location of customized markdown templates (default \"GraphQL Doc Templates\")")
	rootCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "w", true, "Enable overwriting existing documentation")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "out-dir", "o", filepath.Join(must(os.Getwd()).(string), "graphqldocs"), "Specifies the the output directory")
	rootCmd.PersistentFlags().BoolVarP(&format, "format", "f", true, "If --format is present, generator will format graphql schema using prettier")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "If --dry-run is present, generator will run but print generated documentation output into the console.")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
