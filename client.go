package main

import (
	"fmt"
	"net"
	"os"
)

const (
	maxUDPPacketSize = 1400 // Set to a reasonable size less than the typical MTU of 1500
)

func clientmain() {

	// Start listening on port 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", clientListenPort))
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on port %s\n", clientListenPort)

	// Connect to the server on port 8080
	outconn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error Connecting to TCP server:", err)
		os.Exit(1)
	}
	defer outconn.Close()
	fmt.Printf("Connect to %s\n", serverAddress)

	for {
		// Accept an incoming connection
		inconn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleClientConnection(inconn, outconn)
	}

}

func handleClientConnection(inconn net.Conn, outconn net.Conn) {

	recivedData := make([]byte, 1600)

	_, err := inconn.Read(recivedData)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Encrypte data
	encryptedData, err := EncryptECB(recivedData)
	if err != nil {
		fmt.Println("Error while encrypting data:", err)
		return
	}

	// Send a message to the server
	_, err = outconn.Write(encryptedData)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

}

func clientmessage() {
	// Connect to the server on port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "Hello, Server!\ns"

	// Send a message to the server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
}
