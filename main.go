package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	var costs [3]int
	costs[0] = len(primary)
	costs[1] = costs[0] + len(secondary)
	costs[2] = costs[1] + len(tertiary)
	messages := [3]string{primary, secondary, tertiary}
	return messages, costs
}
