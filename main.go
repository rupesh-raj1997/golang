package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	counts := make(map[string]int)

	for _, message := range messages {
		lower := strings.Fields(strings.ToLower(strings.ReplaceAll(message, "!", "")))
		fmt.Println(lower)
		for _, str := range lower {
			if count, ok := counts[str]; ok {
				counts[str] = count + 1
			} else {
				counts[str] = 1
			}
		}
	}

	return len(counts)
}
