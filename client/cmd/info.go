/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"dir/proto/directoryInfo"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
	"os"
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
		}

		//path := args[0]
		//
		//if len(args) == 0 {
		//	cmd.PrintErr(errors.New("You need to specify the path to the file"))
		//}
		//
		//response, err := client.InfoDir(context.Background(), &directoryInfo.PathRequest{Path: path})
		//if err != nil {
		//	grpclog.Fatalf("fail to dial: %v", err)
		//}
		//
		//printDirInfo(response)

	},
}

//func infoInternalMode(client directoryInfo.InfoDirectoryClient) {
//
//	waitc := make(chan string)
//	closed := make(chan bool)
//	var path string
//	//response := make(chan directoryInfo.DirInfoStreamClientResponse)
//	//var path string
//
//	stream, err := client.InfoDirStreamAll(context.Background())
//	if err != nil {
//		grpclog.Fatalf(err.Error())
//	}
//
//	go func() {
//
//		//	in, err := stream.Recv()
//		//	if err == io.EOF {
//		//		close(waitc)
//		//		return
//		//	}
//		//	if err != nil {
//		//		log.Fatalf("client.RouteChat failed: %v", err)
//		//	}
//		//	printDirInfo(in)
//		//}
//	}()
//
//	fmt.Scan(os.Stdin, &path)
//	waitc <- path
//
//	for value := range closed {
//		if value {
//			stream.CloseSend()
//		}
//	}
//
//}

func infoInternalMode(client directoryInfo.InfoDirectoryClient) {

	//response := make(chan directoryInfo.DirInfoStreamClientResponse)
	//var path string

	stream, err := client.InfoDirStreamAll(context.Background())
	if err != nil {
		grpclog.Fatalf(err.Error())
	}

	closed := make(chan bool)
	console := make(chan string)

	//go func() {
	//	var path string
	//	fmt.Scan(os.Stdin, &path)
	//	if path != "" {
	//		consol <- path
	//	} else {
	//		fmt.Println("tut")
	//		closed <- true
	//	}
	//}()
	read := make(chan *directoryInfo.DirInfoResponse)
	go func() {
		for {
			in, _ := stream.Recv()
			read <- in
		}
	}()

	go func() {
		var path string

		for {
			fmt.Scan(os.Stdin, &path)
			if path != "" {
				console <- path
			}
			time.Sleep(time.Second)
		}

	}()

	var path string
	for {
		select {
		case response := <-read:
			fmt.Println(response)
		case <-closed:
			stream.CloseSend()
			return
		case path = <-console:
			stream.Send(&directoryInfo.PathRequest{Path: path})
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
