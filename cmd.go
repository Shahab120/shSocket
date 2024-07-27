package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	key                 string
	serverListenPort    string
	clientServerAddress string
	clientListenPort    string
)

var rootCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	serverCmd.PersistentFlags().StringVar(&key, "key", "", "key")
	serverCmd.PersistentFlags().StringVar(&serverListenPort, "port", "", "listen port")
	rootCmd.AddCommand(serverCmd)

	clientCmd.PersistentFlags().StringVar(&key, "key", "", "key")
	clientCmd.PersistentFlags().StringVar(&clientServerAddress, "addr", "", "server address (host:port)")
	clientCmd.PersistentFlags().StringVar(&serverListenPort, "port", "", "listen port")
	rootCmd.AddCommand(clientCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run an instance of the server",
	Long:  "",
	Run:   runServer,
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run an instance of the client",
	Long:  "",
	Run:   runClient,
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("Run Server")

	servermain()

}

func runClient(cmd *cobra.Command, args []string) {
	fmt.Println("Run Client")

	clientmain()
}
