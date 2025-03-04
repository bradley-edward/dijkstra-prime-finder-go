package main

import (
	"dijkstra-prime-finder/priorityqueue"
)

func dijkstraAlgo(inputValue uint) *[]uint {
	primesList := []uint{2}

	multiplesPq := priorityqueue.New()
	multiplesPq.Enqueue(2, 4)

	var num uint
	for num = 3; num <= inputValue; num++ {
		// Find the smallest "multiple" in the PQ.
		minMultiple := multiplesPq.Peek()
		// If the current number is less than the smallest multiple, the current number is a prime.
		if num < minMultiple.Priority {
			primesList = append(primesList, num)
			multiplesPq.Enqueue(num, num*num)
			continue
		}

		// Else, the number is NOT a prime.
		for num == minMultiple.Priority {
			multiplesPq.OverwriteMinItem(minMultiple.Value, minMultiple.Value+minMultiple.Priority)
			minMultiple = multiplesPq.Peek()
		}
	}

	return &primesList
}