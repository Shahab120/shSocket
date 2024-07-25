package main

import (
	"fmt"
	"net"
)

func clientmain() {
	// Resolve the server address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// Create the UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error creating connection:", err)
		return
	}
	defer conn.Close()

	// Message to send to the server
	message := "Hello, UDP server!"

	// Write the message to the server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to UDP connection:", err)
		return
	}

	// Buffer to hold the server's response
	buf := make([]byte, 1024)

	// Read the response from the server
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from UDP connection:", err)
		return
	}

	// Print the server's response
	fmt.Printf("Received response from server: %s\n", string(buf[:n]))
}
