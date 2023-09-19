/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dir/proto/directoryInfo"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
	"io"
	"time"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "directory info",
	Long:  `directory info`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := getConnection(cmd)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		client := directoryInfo.NewInfoDirectoryClient(conn)

		internal, _ := cmd.Flags().GetBool("internal")
		if internal {
			infoInternalMode(client)
			return
		}

		if len(args) == 0 {
			cmd.PrintErr(errors.New("You need to specify the path to the file\n"))
			return
		}
		path := args[0]

		response, err := client.InfoDir(context.Background(), &directoryInfo.PathRequest{Path: path})
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		printDirInfo(response)

	},
}

func infoInternalMode(client directoryInfo.InfoDirectoryClient) {

	exit := false

	stream, err := client.InfoDirStreamAll(context.Background())
	if err != nil {
		grpclog.Fatalf(err.Error())
	}

	go func() {
		for {
			if exit {
				return
			}
			in, err := stream.Recv()
			if err == io.EOF {
				continue
			}
			if err != nil {
				grpclog.Fatalf(err.Error())
			}
			printDirInfo(in)
		}
	}()

	var path string

	for {
		fmt.Println("input Directory")
		fmt.Scanf("%s\n", &path)
		if path == "exit" {
			stream.CloseSend()
			exit = true
			return
		}
		if path != "" {
			stream.Send(&directoryInfo.PathRequest{Path: path})
			time.Sleep(time.Millisecond)
		}
	}

}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().String("conn", "127.0.0.1:5300", "connection grpc server")
	infoCmd.Flags().BoolP("internal", "i", false, "internal")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
