package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, m := range msg {
		for _, bword := range badWords {
			if m == bword {
				return i
			}
		}
	}
	return -1
}
