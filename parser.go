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

func parseQuestion(packet []byte, offset int) (DNSQuestion, int, error) {

	//Read QNames
	Qname := ""
	for {
		length := int(packet[offset])
		offset++
		if length == 0 {
			break
		}
		Qname += string(packet[offset : offset+length])
		Qname += "."
		offset += length
	}
	Qname = Qname[0 : len(Qname)-1]

	//Read QType
	Qtype := binary.BigEndian.Uint16(packet[offset : offset+2])
	offset += 2

	//Read QClass
	Qclass := binary.BigEndian.Uint16(packet[offset : offset+2])
	offset += 2

	return DNSQuestion{Qname, Qtype, Qclass}, offset, nil
}

func resolveQuestion(question DNSQuestion) []DNSResourceRecord {
	answer := DNSResourceRecord{
		Name:  question.QName,
		Type:  question.QType,
		Class: question.QClass,
		TTL:   60,
		RData: []byte{8, 8, 8, 8},
	}

	return []DNSResourceRecord{answer}
}
