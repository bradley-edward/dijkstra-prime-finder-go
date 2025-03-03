package main

import "math"

func eratosthenesSieve(inputValue uint) *[]uint {
	primesList := []uint{}
	boolArray := make([]bool, inputValue+1)

	var num, dividedInput, multiplier uint
	var roundedSquareRoot uint = uint(math.Floor(math.Sqrt(float64(inputValue))))

	for num = 1; num <= inputValue; num++ {
		boolArray[num] = true
	}
	boolArray[1] = false

	for num = 2; num <= roundedSquareRoot; num++ {
		if !boolArray[num] {
			continue
		}

		primesList = append(primesList, num)
		dividedInput = inputValue / num

		for multiplier = 2; multiplier <= dividedInput; multiplier++ {
			boolArray[num*multiplier] = false
		}
	}
	for num = roundedSquareRoot + 1; num <= inputValue; num++ {
		if boolArray[num] {
			primesList = append(primesList, num)
		}
	}

	return &primesList
}
