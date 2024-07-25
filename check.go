package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func check() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	for _, device := range devices {
		fmt.Printf("Name: %s\n", device.Name)
		fmt.Printf("Description: %s\n", device.Description)
		for _, address := range device.Addresses {
			fmt.Printf("IP address: %s\n", address.IP)
			fmt.Printf("Subnet mask: %s\n", address.Netmask)
		}
		fmt.Println()
	}
}
