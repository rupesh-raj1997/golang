package main

func getMessageCosts(messages []string) []float64 {
	size := len(messages)
	costs := make([]float64, size, size)
	for i, message := range messages {
		costs[i] = 0.01 * float64(len(message))
	}
	return costs
}
