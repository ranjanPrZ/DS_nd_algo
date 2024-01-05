// write a program to sort an array using bubble sort algorithm in Go language.

package main

import "fmt"

func bubble_fort(arr []int){
	arr_length := len(arr)
	for i := 0; i < arr_length-1; i++ {
			swapped := false
			for j := 0; j < arr_length-i-1; j++ {
					if arr[j] > arr[j+1] {
							arr[j], arr[j+1] = arr[j+1], arr[j]
							swapped = true
					}
			}
			if !swapped {
					break
			}
	}
}


func main() {
	arr := []int{40, 50, 20, 0, -10, 30, 10}
	bubble_fort(arr)
	fmt.Println(arr) //Output: [-10 0 10 20 30 40 50]
}