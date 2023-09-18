/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dir/proto/directoryInfo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
)

// testStreamCmd represents the testStream command
var testStreamCmd = &cobra.Command{
	Use:   "test-stream",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getConnection(cmd)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		client := directoryInfo.NewInfoDirectoryClient(conn)
		stream, err := client.InfoDirStreamClient(context.Background())
		if err != nil {
			grpclog.Fatal(err)
		}

		var pathReqs []*directoryInfo.PathRequest

		for _, arg := range args {
			pathReqs = append(pathReqs, &directoryInfo.PathRequest{Path: arg})
		}

		for _, pathReq := range pathReqs {
			if err := stream.Send(pathReq); err != nil {
				grpclog.Fatal(err)
			}
		}

		replay, err := stream.CloseAndRecv()
		if err != nil {
			grpclog.Fatal(err)
		}

		for _, value := range replay.GetDirectoriesInfo() {
			printDirInfo(value)
		}

	},
}

func init() {
	rootCmd.AddCommand(testStreamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testStreamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testStreamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
