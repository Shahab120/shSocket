package main

import (
	"fmt"
	"net"
)

func servermain() {
	// Resolve the address
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create the UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating connection:", err)
		return
	}
	defer conn.Close()

	// Buffer to hold incoming data
	buf := make([]byte, 1024)

	for {
		// Read data from the connection
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}

		// Print the received message
		fmt.Printf("Received message from %s: %s\n", clientAddr, string(buf[:n]))

		// Echo the message back to the client
		_, err = conn.WriteToUDP(buf[:n], clientAddr)
		if err != nil {
			fmt.Println("Error writing to UDP connection:", err)
			continue
		}
	}
}
