package main

import "fmt"

var refArr []int = []int{5,4,7,2,9,3,6}

func bubbleSort(arr []int) []int {
	for j := 1; j<len(arr); j++ {
		for i:=0; i<len(arr)-j;i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}

func main() {
	fmt.Println(bubbleSort(refArr))
}
