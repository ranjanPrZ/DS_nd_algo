package main

import (
	"fmt"
	"math"
)

func radixSort(arr []int) []int {
	maxDigitCount := maxDigits(arr)

	for digitPlace := 0; digitPlace < maxDigitCount; digitPlace++ {
		buckets := make([][]int, 10)

		for _, num := range arr {
			digit := getDigit(num, digitPlace)
			buckets[digit] = append(buckets[digit], num)
		}

		arr = flatten(buckets)
	}

	return arr
}

func maxDigits(arr []int) int {
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return int(math.Floor(math.Log10(float64(max)))) + 1
}

func getDigit(num, place int) int {
	return (num / int(math.Pow(10, float64(place)))) % 10
}

func flatten(buckets [][]int) []int {
	var result []int

	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

func main() {
	inputArray := []int{170, 45, 75, 90, 802, 24, 2, 66}
	sortedArray := radixSort(inputArray)

	fmt.Printf("Original array: %v\n", inputArray)
	fmt.Printf("Sorted array: %v\n", sortedArray)
}
