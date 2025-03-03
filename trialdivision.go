package main

import "math"

func trialDivisionPrimeCheck(inputNum uint, primesList *[]uint) bool {
	var roundedSquareRoot uint = uint(math.Floor(math.Sqrt(float64(inputNum))))
	for primeIdx := 0; (*primesList)[primeIdx] <= roundedSquareRoot; primeIdx++ {
		if inputNum % (*primesList)[primeIdx] == 0 {
			return false
		}
	}

	return true
}

func trialDivision(inputValue uint) *[]uint {
	primesList := []uint{2,3}

	var num uint

	for num = 4; num <= inputValue; num++ {
		if trialDivisionPrimeCheck(num, &primesList) {
			primesList = append(primesList, num)
		}
	}

	return &primesList
}