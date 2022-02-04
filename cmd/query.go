/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
    "github.com/peng225/wolper/service"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "hogehoge",
	Long: `fugafuga`,
	Run: func(cmd *cobra.Command, args []string) {
        address, err := cmd.Flags().GetString("address")
        if err != nil {
            panic(err)
        }
        // TODO: address format check
        fmt.Println("address:", address)

        port, err := cmd.Flags().GetInt("port")
        if err != nil {
            panic(err)
        } else if port <= 0 {
            panic("The port must be a positive number.")
        }
        fmt.Println("port:", port)

        include, err := cmd.Flags().GetString("include")
        if err != nil {
            panic(err)
        }
        fmt.Println("include:", include)

        exclude, err := cmd.Flags().GetString("exclude")
        if err != nil {
            panic(err)
        }
        fmt.Println("exclude:", exclude)

        fixed, err := cmd.Flags().GetString("fixed")
        if err != nil {
            panic(err)
        }
        fmt.Println("fixed:", fixed)

        service.ClientQuery(address + ":" + strconv.Itoa(port),
            include, exclude, fixed)
    },
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	queryCmd.Flags().StringP("address", "a", "localhost", "The IP address to connect")
	queryCmd.Flags().IntP("port", "p", 8080, "The port number to connect")
	queryCmd.Flags().StringP("include", "i", "", "Included characters (can contain duplicated characters)")
	queryCmd.Flags().StringP("exclude", "e", "", "Excluded characters (can contain duplicated characters)")
	queryCmd.Flags().StringP("fixed", "f", "", "Fiexed pattern (eg. \"sp...\" can match strings like \"spawn\", \"speak\", \"spray\", and so on)")
}
