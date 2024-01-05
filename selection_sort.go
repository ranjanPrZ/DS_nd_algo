package main

import "fmt"

func main() {
    n := []int{1, 39, 2, 9, 7, 54, 11}
    
    for i := 0; i < len(n)-1; i++ {
        minIndex := i

        // Find the index of the minimum element in the unsorted part of the array
        for j := i + 1; j < len(n); j++ {
            if n[j] < n[minIndex] {
                minIndex = j
            }
        }

        // Swap the found minimum element with the first element of the unsorted part
        if minIndex != i {
            n[i], n[minIndex] = n[minIndex], n[i]
        }
    }

    fmt.Println(n)
}

