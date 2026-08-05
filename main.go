package main

import (
	"fmt"
	"net"
)

func main() {

	addr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 53,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("DNS Server listening on port 53")

	buffer := make([]byte, 512)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Received packet from:", clientAddr)
		fmt.Println("Packet length:", n)
		fmt.Println(buffer[:n])
	}
}
