package main

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		for j := i + 1; j < groupSize; j++ {
			connections += 1
		}
	}
	return connections
}
