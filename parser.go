package main

import (
	"encoding/binary"
	"errors"
)

func parseHeader(packet []byte) (DNSHeader, error) {
	if len(packet) < 12 {
		return DNSHeader{}, errors.New("Packet length too small")
	}
	id := binary.BigEndian.Uint16(packet[0:2])
	flags := binary.BigEndian.Uint16(packet[2:4])
	qdcount := binary.BigEndian.Uint16(packet[4:6])
	ancount := binary.BigEndian.Uint16(packet[6:8])
	nscount := binary.BigEndian.Uint16(packet[8:10])
	arcount := binary.BigEndian.Uint16(packet[10:12])
	return DNSHeader{id, flags, qdcount, ancount, nscount, arcount}, nil
}
