package main

import (
	"log"
	"math"
	"reflect"
)

// 正方形
type rectangle struct {
	width, height float64
}

type circle struct {
	r float64  // 半径
}

// 计算正方形面积
func (r rectangle)area() float64 {
	return r.width * r.height
}

// 计算圆面积
func (c circle)area() float64 {
	return math.Pi * c.r * c.r
}

//计算正方形周长
func (r rectangle)zhouchang() float64 {
	return 2 * (r.width + r.height)
}

// 计算圆周长
func (c circle) cZhouchang() float64 {
	return 2 * math.Pi * c.r
}

// main函数
func main() {
	// 定义一个正方形
	newRectangle := rectangle{10, 5}

	area := newRectangle.area()
	log.Println("正方形的面积：", area)
	log.Println("area的类型:", reflect.TypeOf(area))

	zhouchang := newRectangle.zhouchang()
	log.Println("正方形的周长：", zhouchang)
	log.Println("zhouchang的类型: ", reflect.TypeOf(zhouchang))

	//定义一个圆
	newCircle := circle{5}

	carea := newCircle.area()
	log.Println("圆的面积: ", carea)
	log.Println("carea的类型: ", reflect.TypeOf(carea))

	czhouchang := newCircle.cZhouchang();
	log.Println("圆的周长: ", czhouchang)
	log.Println("czhouchagn 的类型: ", reflect.TypeOf(czhouchang))

}
