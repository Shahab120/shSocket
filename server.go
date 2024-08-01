package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func servermain() {
	// Resolve the address to listen on
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%s", serverListenPort))
	if err != nil {
		log.Fatal(err)
	}

	// Set up the UDP listener
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Resolve the forward address
	forwardAddr, err := net.ResolveUDPAddr("udp", serverForwardAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the UDP connection to forward packets
	forwardConn, err := net.DialUDP("udp", nil, forwardAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer forwardConn.Close()

	buf := make([]byte, 1600)

	for {
		// Read data from the connection
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading from UDP connection:", err)
			continue
		}
		decreptedPacket, err := DecryptECB(buf[:n])
		if err != nil {
			log.Println("Error reading from UDP connection:", err)
			continue
		}

		// Parse the packet
		packet := gopacket.NewPacket(decreptedPacket, layers.LayerTypeEthernet, gopacket.Default)
		fmt.Printf("Received packet from : %v\n", packet.String())
		fmt.Printf("Payload %s", string(packet.Data()))

		_, err = forwardConn.Write(packet.Data())
		if err != nil {
			log.Println("Error forwarding packet:", err)
		}
	}
}
