package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func clientmain() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <device> <server:port>\n", os.Args[0])
		os.Exit(1)
	}

	device := os.Args[1]
	serverAddr := os.Args[2]

	// Open the device for capturing
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up UDP connection to the server
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Start capturing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Serialize the packet
		packetData := packet.Data()

		// Send the serialized packet to the server
		_, err := conn.Write(packetData)
		if err != nil {
			log.Println("Error sending packet:", err)
		}
	}
}
