/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"

	"github.com/peng225/wolper/service"

	"github.com/peng225/wolper/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run a server for words querying.",
	Long: `Run a server for words querying.
This sub-command should be executed before query sub command is used.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			panic(err)
		} else if input == "" {
			panic("The input must not be empty.")
		}
		fmt.Println("input:", input)

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		} else if port <= 0 {
			panic("The port must be a positive number.")
		}
		fmt.Println("port:", port)

		listenPort, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
		if err != nil {
			fmt.Println("failed to listen:", err)
		}

		registrar := grpc.NewServer()
		wssi := service.WolperServiceServerImpl{}
		wssi.Init(input)
		pb.RegisterWolperServiceServer(registrar, &wssi)
		reflection.Register(registrar)
		fmt.Println("Service started.")
		registrar.Serve(listenPort)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().IntP("port", "p", 8081, "The port number to be opend")
	serverCmd.Flags().StringP("input", "i", "dict.txt", "The path to the dictionary file")
}
