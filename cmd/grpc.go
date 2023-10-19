/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MatheusBenetti/desafio1/application/grpc"
	"github.com/MatheusBenetti/desafio1/infra/db"
	"github.com/spf13/cobra"
	"os"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectToDB(os.Getenv("ENV"))
		grpc.StartGrpcServer(database, 50051)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "grpc server port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
