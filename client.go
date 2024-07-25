package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func clientmain() {

	device := `\Device\NPF_Loopback`
	serverAddr := "192.168.1.5:8080"
	port := "8182"

	// Check if the device exists
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(devices)

	deviceExists := false
	for _, dev := range devices {
		if dev.Name == device {
			deviceExists = true
			break
		}
	}

	if !deviceExists {
		log.Fatalf("Device %s not found", device)
	}

	// Open the device for capturing
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up a BPF filter for the specified port
	filter := fmt.Sprintf("port %s", port)
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal("Error setting BPF filter: ", err)
	}
	fmt.Println("Only capturing packets on port", port)

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
		fmt.Println(packet)
		fmt.Println(string(packet.Data()))
		// Send the serialized packet to the server
		_, err := conn.Write(packetData)
		if err != nil {
			log.Println("Error sending packet:", err)
		}
	}

}

func clientmessage() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8182")
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
	message := "Hello, UDP mserver!"

	// Write the message to the server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to UDP connection:", err)
		return
	}

}
