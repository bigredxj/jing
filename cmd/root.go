package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "jing",
		Short: "A generator for Cobra based Applications",
		Long: `jing is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		TraverseChildren: true,
		Version:          "9.9.9",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "default license", "name of license for the project")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
