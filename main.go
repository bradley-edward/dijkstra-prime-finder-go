package main

import (
	"fmt"
)

func main() {
	var inputValue uint = 1500
	if inputValue < 3 {
		return
	}

	//eratosthenesSieve(inputValue)

	//primesList := eratosthenesSieve(inputValue)//dijkstraAlgo(inputValue)

	fmt.Println(*(dijkstraAlgo(inputValue)))
	//fmt.Println(*(eratosthenesSieve(inputValue)))
	//fmt.Println(*(trialDivision(inputValue)))
	//fmt.Println("Primes Count:", len(*primesList))
}