/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/peng225/wolper/web"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server for words querying.",
	Long:  `Run a web server for words querying.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		} else if port <= 0 {
			panic("The port must be a positive number.")
		}
		fmt.Println("port:", port)

		wolper_port, err := cmd.Flags().GetInt("wolper_port")
		if err != nil {
			panic(err)
		} else if wolper_port <= 0 {
			panic("The wolper_port must be a positive number.")
		}
		fmt.Println("wolper_port:", wolper_port)

		address, err := cmd.Flags().GetString("address")
		if err != nil {
			panic(err)
		}
		// TODO: address format check
		fmt.Println("address:", address)

		web.Start(port, address+":"+strconv.Itoa(wolper_port), "web")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webCmd.Flags().IntP("port", "p", 8080, "The port number to be opend")
	webCmd.Flags().Int("wolper_port", 8081, "The port number of the wolper server")
	webCmd.Flags().StringP("address", "a", "localhost", "The IP address of the wolper server")
}
