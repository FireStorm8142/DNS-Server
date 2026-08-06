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

	//Main server loop
	buffer := make([]byte, 512)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println()
		fmt.Println("Received packet from:", clientAddr)
		fmt.Println("Packet length:", n)

		//Parse Header
		header, err := parseHeader(buffer[:n])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%+v\n", header)

		//Parse QuestionSection
		index := 12
		questions := make([]DNSQuestion, int(header.QDCount))
		for i := 0; i < int(header.QDCount); i++ {
			question, offset, err := parseQuestion(buffer[:n], index)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%+v\n", question)
			questions[i] = question
			index = offset
		}

		//Build Packet
		packet := DNSPacket{header, questions, nil, nil, nil}

		response := buildResponse(packet)
	}
}
