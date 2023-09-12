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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
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

		client := directoryInfo.NewInfoDirectoryClient(conn)

		if len(args) == 0 {
			log.Fatal(errors.New("You need to specify the path to the file"))
		}

		response, err := client.InfoDir(context.Background(), &directoryInfo.PathRequest{Path: args[0]})
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		printDirInfo(response)

	},
}

func getConnection(cmd *cobra.Command) (*grpc.ClientConn, error) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	url, err := cmd.Flags().GetString("conn")
	if err != nil {
		return nil, err
	}

	if url == "" {
		return nil, errors.New("conn cannot be empty")
	}
	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func printDirInfo(info *directoryInfo.DirInfoResponse) {

	files := info.Files
	size := info.Size
	directories := info.Directories

	fmt.Println(fmt.Sprintf("sizeAll %d\n", size))
	fmt.Println("directories:")
	for _, dir := range directories {
		fmt.Println(fmt.Sprintf("name: %s, size (byte): %d", dir.Name, dir.Size))
	}
	fmt.Println()
	fmt.Println("files:")
	for _, file := range files {
		fmt.Println(fmt.Sprintf("name: %s, size (byte): %d", file.Name, file.Size))
	}

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
