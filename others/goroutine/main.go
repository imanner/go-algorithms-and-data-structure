package main

import "log"

type people struct {
	name string
	sex string
	age int
}

var nchan = make(chan int, 1)

func (p people)say() {
	for i:=0; i<10; i++ {
		log.Println("my name is: ", p.name)
		log.Println("my sex is: ", p.sex)
		log.Println("my age is: ", p.age)
	}

	nchan <- 1000
}

func main() {
	me := people{"zhang", "male", 20}
	go me.say()

	for {
		select {
		case <- nchan:
			log.Println("输出完毕")
		}
	}
}
