package main

func buildResponse(packet DNSPacket) []byte {
	var response []byte

	//Build Header
	response = append(response, buildHeader(packet.Header))

	//Build QuestionSection
	for _, question := range packet.Questions {
		response = append(response, buildQuestionSection(question))
	}

	return response
}
