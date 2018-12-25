package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
整体思路
 基线条件：如果数组的长度小于等于1 { 返回只有一个元素的数组 }
 获取标准值：数组中随机抽取一个值
 生成3个容器：低段， 中断， 高段，长度初始为0，容量初始为原数组的长度
 遍历元素组：把所有数据分别放到3个容器中构成整体有序的
 针对低段和高段数组再次进行快速排序
 拼合3段：以低段为基础分别拼合，这是最大的难点
 返回最终拼合结果
*/
func quickSort(arr []int) []int {
	// 保证是数组而不是一个值
	if len(arr) <= 1 {
		return arr
	}

	// 随机去一个中间值
	mideItem := arr[rand.Intn(len(arr))]
	lowPart := make([]int, 0, len(arr))
	midemPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))

	//遍历arr 将不同的部分放到不同的区域
	for _, item := range arr {
		if item > mideItem {
			highPart = append(highPart, item)
		} else if item == mideItem {
			midemPart = append(midemPart, item)
		} else if item < mideItem {
			lowPart = append(lowPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, midemPart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

// 构建随机的数组
func makeRandArr(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func main() {
	initArr := makeRandArr(10)
	fmt.Println("初始化的数组是: ", initArr)
	res := quickSort(initArr)
	fmt.Println("排序后的结果是: ", res)
}
