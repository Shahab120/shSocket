package main

import (
	"fmt"
	"net"
	"os"
)

func servermain() {

	// Start listening on port 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serverListenPort))
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on port %s\n", serverListenPort)

	// Connect to the server on port 8080
	outconn, err := net.Dial("tcp", serverForwardAddress)
	if err != nil {
		fmt.Println("Error Connecting to TCP server:", err)
		os.Exit(1)
	}
	defer outconn.Close()
	fmt.Printf("Connect to %s\n", serverForwardAddress)

	for {
		// Accept an incoming connection
		inconn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleServerConnection(inconn, outconn)
	}

}

func handleServerConnection(inconn net.Conn, outconn net.Conn) {

	recivedData := make([]byte, 1600)

	// Read data from the connection
	_, err := inconn.Read(recivedData)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	fmt.Println("Recived: ", string(recivedData))

	// Encrypte data
	decryptedData, err := DecryptECB(recivedData)
	if err != nil {
		fmt.Println("Error while decrypting data:", err)
		return
	}

	// Send a message to the server
	_, err = outconn.Write(decryptedData)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

}
