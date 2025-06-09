package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		firstCh := runes[0]

		if savedMap, ok := nameCounts[firstCh]; ok {
			if count, present := savedMap[name]; present {
				savedMap[name] = count + 1
			} else {
				savedMap[name] = 1
			}
		} else {
			nameCounts[firstCh] = map[string]int{name: 1}
		}
	}

	return nameCounts
}
