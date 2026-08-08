package main

func buildResponse(packet DNSPacket) []byte {
	var response []byte

	//Build Header
	response = append(response, buildHeader(packet.Header, uint16(len(packet.Answers)))...)
	//Build QuestionSection
	for _, question := range packet.Questions {
		response = append(response, buildQuestionSection(question)...)
	}
	//Build Answer
	for _, answer := range packet.Answers {
		response = append(response, buildAnswer(answer)...)
	}
	return response
}
