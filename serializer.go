package main

import (
	"encoding/binary"
	"strings"
)

func buildHeader(header DNSHeader, ansCount uint16) []byte {
	response := make([]byte, 12)
	binary.BigEndian.PutUint16(response[0:2], header.ID)
	flags := header.Flags
	flags = flags | 1<<15
	flags &^= 1 << 7
	flags &^= 0x000F
	binary.BigEndian.PutUint16(response[2:4], flags)
	binary.BigEndian.PutUint16(response[4:6], header.QDCount)
	binary.BigEndian.PutUint16(response[6:8], ansCount)
	binary.BigEndian.PutUint16(response[8:10], 0)
	binary.BigEndian.PutUint16(response[10:12], 0)
	return response
}

func buildQuestionSection(question DNSQuestion) []byte {
	var response []byte
	labels := strings.Split(question.QName, ".")
	for _, label := range labels {
		response = append(response, byte(len(label)))
		response = append(response, []byte(label)...)
	}
	response = append(response, byte(0))

	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, question.QType)
	response = append(response, temp...)
	binary.BigEndian.PutUint16(temp, question.QClass)
	response = append(response, temp...)
	return response
}

func buildAnswer(answer DNSResourceRecord) []byte {
	var response []byte
	labels := strings.Split(answer.Name, ".")
	for _, label := range labels {
		response = append(response, byte(len(label)))
		response = append(response, []byte(label)...)
	}
	temp := make([]byte, 4)
	binary.BigEndian.PutUint16(temp[0:2], answer.Type)
	response = append(response, temp[0:2]...)
	binary.BigEndian.PutUint16(temp[0:2], answer.Class)
	response = append(response, temp[0:2]...)
	binary.BigEndian.PutUint32(temp[0:4], answer.TTL)
	response = append(response, temp[0:4]...)
	binary.BigEndian.PutUint16(temp[0:2], uint16(len(answer.RData)))
	response = append(response, temp[0:2]...)
	response = append(response, answer.RData...)
	return response
}

func buildAuthority() []byte {
	return nil
}

func buildAdditional() []byte {
	return nil
}
