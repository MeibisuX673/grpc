/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dir/proto/directoryInfo"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "directory info",
	Long:  `directory info`,
	Run: func(cmd *cobra.Command, args []string) {
		dirInfo, err := getDirInfo(args[0])
		if err != nil {
			cmd.Println(err)
		}

		printDirInfo(dirInfo)

	},
}

func printDirInfo(info *directoryInfo.DirInfoResponse) {

	files := info.Files
	size := info.Size
	directories := info.Directories

	fmt.Println(fmt.Sprintf("sizeAll %d\n", size))
	fmt.Println("directories:")
	for _, dir := range directories {
		fmt.Println(fmt.Sprintf("name: %s, size: %d", dir.Name, dir.Size))
	}
	fmt.Println()
	fmt.Println("files:")
	for _, file := range files {
		fmt.Println(fmt.Sprintf("name: %s, size: %d", file.Name, file.Size))
	}

}

func getDirInfo(path string) (*directoryInfo.DirInfoResponse, error) {
	opts := []grpc.DialOption{}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := directoryInfo.NewInfoDirectoryClient(conn)

	response, err := client.InfoDir(context.Background(), &directoryInfo.PathRequest{Path: path})
	if err != nil {
		log.Fatal(err)
	}

	//log.Fatal(response)

	return response, nil

}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
