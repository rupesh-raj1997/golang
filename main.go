package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		count, ok := validUsers[user]

		if ok {
			validUsers[user] = count + 1
		}
	}
}
