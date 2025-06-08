package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	total := make([]float64, 0)
	for _, cost := range costs {
		if cost.day == day {
			total = append(total, cost.value)
		}
	}
	return total
}
