package main

import "fmt"

var refArr []int = []int{1,6,2,4,9,10}

// 这种方式虽然笨 但是二分查找要求有序，而这种并不要求，并且实现超简单，思维也简单
func linearSearch(arr []int, aim int) int {
	pos := -1
	arrlen := len(arr)

	for i:=0; i<arrlen; i++ {
		if arr[i] == aim {
			pos = i
			break
		}
	}

	return pos
}

func main() {
	res := linearSearch(refArr, 3)

	fmt.Println("找到的结果是:", res)
}
