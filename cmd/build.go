/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/peng225/wolper/dictionary"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a n-character dictionary for querying.",
	Long: `Build a n-word dictionary for querying.
This sub-command should be executed before server sub command is used.`,
	Run: func(cmd *cobra.Command, args []string) {
        input, err := cmd.Flags().GetString("input")
        if err != nil {
            panic(err)
        } else if input == "" {
            panic("The input must not be empty.")
        }
        fmt.Println("input:", input)

        output, err := cmd.Flags().GetString("output")
        if err != nil {
            panic(err)
        } else if input == "" {
            panic("The output must not be empty.")
        }
        fmt.Println("output:", output)

        var dict dictionary.Dictionary
        dict.Build(input, output)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildCmd.Flags().StringP("input", "i", "input", "The path to input directory")
	buildCmd.Flags().StringP("output", "o", "dict.txt", "The output file path")
}
