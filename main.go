package main

import (
	"log"
	"os"

	"grpc_helloworld/client"
	"grpc_helloworld/server"

	"github.com/spf13/cobra"
)

func main() {
	var cmdServer = &cobra.Command{
		Use:   "server",
		Short: "start server",
		Long:  `start server.`,
		Run:   server.StartServer,
	}

	var cmdClient = &cobra.Command{
		Use:   "client",
		Short: "start client",
		Long:  `start client.`,
		Run:   client.StartClient,
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdServer, cmdClient)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed cmd: %v", err)
		os.Exit(-1)
	}
}
