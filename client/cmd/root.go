/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dir/proto/directoryInfo"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

var (
	crtFile = "server.crt"
)

const DEFAULT_CONNECTION = "127.0.0.1:5300"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dir",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("root")
	//},
	//PreRun: func(cmd *cobra.Command, args []string) {
	//	fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	//},
}

func printDirInfo(info *directoryInfo.DirInfoResponse) {

	files := info.Files
	size := info.Size
	directories := info.Directories

	fmt.Println(fmt.Sprintf("%s size (byte): %d\n", info.GetPath(), size))
	fmt.Println("directories:")
	for _, dir := range directories {
		fmt.Println(fmt.Sprintf("name: %s, size (byte): %d", dir.Name, dir.Size))
	}
	fmt.Println()
	fmt.Println("files:")
	for _, file := range files {
		fmt.Println(fmt.Sprintf("name: %s, size (byte): %d", file.Name, file.Size))
	}
	fmt.Println()

}

func getConnection(cmd *cobra.Command) (*grpc.ClientConn, error) {

	url, err := cmd.Flags().GetString("conn")
	if err != nil {
		url = DEFAULT_CONNECTION
	}
	if url == "" {
		return nil, errors.New("conn cannot be empty")
	}

	creds, err := credentials.NewClientTLSFromFile(crtFile, DEFAULT_CONNECTION)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
