package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	maxUDPPacketSize = 1400 // Set to a reasonable size less than the typical MTU of 1500
)

func clientmain() {

	device := `\Device\NPF_Loopback`
	serverAddr := "localhost:8080"
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

		// Split packet data if it exceeds the maximum UDP packet size
		for len(packetData) > 0 {
			chunkSize := maxUDPPacketSize
			if len(packetData) < chunkSize {
				chunkSize = len(packetData)
			}
			chunk := packetData[:chunkSize]
			packetData = packetData[chunkSize:]

			encryptedData, err := EncryptECB(chunk)
			if err != nil {
				log.Println("Error sending packet:", err)
				continue
			}

			// gopacket.NewLayerClass()

			// Send the chunk to the server4
			_, err = conn.Write(encryptedData)
			if err != nil {
				log.Println("Error sending packet:", err)
			}
		}
	}

}

func clientmessage() {
	serverAddr, err := net.ResolveUDPAddr("udp", "192.168.1.8:8182")
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
	// message, err := os.ReadFile("F:/shSocket/base64.txt")
	// if err != nil {
	// 	fmt.Println("Error creating connection:", err)
	// 	return
	// }

	message := []byte("hi there")

	// Write the message to the server
	for {
		chmessage := message
		for len(chmessage) > 0 {
			chunkSize := maxUDPPacketSize
			if len(chmessage) < chunkSize {
				chunkSize = len(chmessage)
			}
			chunk := chmessage[:chunkSize]
			chmessage = chmessage[chunkSize:]

			// Send the chunk to the server
			_, err := conn.Write(chunk)
			if err != nil {
				log.Println("Error sending packet:", err)
			}
		}
		time.Sleep(1 * time.Second)
	}

}
