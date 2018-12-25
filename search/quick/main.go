package main

import (
	"fmt"
)

/*
整体思路
 数组首先必须是有序的！这个一定要注意！
 基线条件：如果左侧的游标已经和右侧游标相遇就应该停下
 获取中间位置：(右侧-左侧)/2
 用数组的中间值和目标值对比：
	如果中间值小于目标值，说明目标值在右半部分：
		左侧游标位置更新
		针对右半部分进行快速查找
	如果中间值大于目标值，说明在左侧：
		右侧游标位置更新
		针对左半部分快速查找
	如果恰好相等，找到了目标值
		返回下标
	最终如果没有找到，返回-1 ，-1标识一个不存在的位置
*/
func quickSearch(arr []int, num int, leftPos int, rightPos int) int {
	pos := -1

	// 保证是数组而不是一个值
	if leftPos >= rightPos {
		return pos
	}

	// 取到中间位置
	midePos 	:= (rightPos - leftPos)/2

	if num < arr[midePos] {
		rightPos = midePos-1
		pos = quickSearch(arr, num, leftPos, rightPos)
	} else if num > arr[midePos] {
		leftPos = midePos + 1
		pos = quickSearch(arr, num, leftPos, rightPos)
	} else {
		pos = midePos
	}

	return pos
}

// 包装款速查询方法
func qs(initArr []int, aim int) int {
	res := quickSearch(initArr, aim, 0, len(initArr))
	return res
}

func main() {
	//initArr := makeRandArr(10)
	initArr := []int{1,3,5,7,9,10,23}
	fmt.Println("初始化的数组是: ", initArr)
	res := qs(initArr, 3)
	fmt.Println("查找的结果是: ", res)
}
