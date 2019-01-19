package main

import "log"

type people struct {
	name string
	age int
	sex string
}

type student struct {
	people // 引入
	xuehao string
}

func main() {
	xiaoMing := student{people{"小明", 10, "男"}, "12312312"}

	log.Println(xiaoMing)
	log.Println(xiaoMing.name)
	log.Println(xiaoMing.age)
	log.Println(xiaoMing.sex)
	log.Println(xiaoMing.xuehao)
}
