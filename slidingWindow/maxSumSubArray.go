package main

import "fmt"

/*
func maxSum(Arr []int, K int) int {
	windowSum := 0
	maximumSum := 0

	for i := 0; i < len(Arr); i++ {
		windowSum += Arr[i]

		if i >= K-1 {
			if maximumSum < windowSum {
				maximumSum = windowSum
			}
			windowSum -= Arr[i-(K-1)]
		}
	}
	return maximumSum
}
*/

//BruteForce method..

func maxSum(Arr []int, K int) int {

	windowSum := 0
	maximumSum := 0

	for i := 0; i <= len(Arr)-K; i++ {

		for j := i; j < i+K; j++ {

			windowSum += Arr[j]

		}
		if maximumSum < windowSum {
			maximumSum = windowSum
			windowSum = 0
		}
	}
	return maximumSum
}

func main() {
	fmt.Println(maxSum([]int{100, 200, 300, 400}, 2))
}
