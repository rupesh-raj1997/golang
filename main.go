package main

func maxMessages(thresh int) int {
	var maxMessages = 0

	for i := 0; ; i++ {
		msgCost := 100 + i

		if thresh >= msgCost {
			thresh -= msgCost
			maxMessages += 1
		} else {
			break
		}
	}

	return maxMessages
}
