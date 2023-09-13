/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dir/proto/directoryInfo"
	"errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "directory info",
	Long:  `directory info`,
	Run: func(cmd *cobra.Command, args []string) {

		path := args[0]
		conn, err := getConnection(cmd)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		client := directoryInfo.NewInfoDirectoryClient(conn)

		if len(args) == 0 {
			cmd.PrintErr(errors.New("You need to specify the path to the file"))
		}

		response, err := client.InfoDir(context.Background(), &directoryInfo.PathRequest{Path: path})
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		printDirInfo(response)

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().String("conn", "127.0.0.1:5300", "connection grpc server")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
