package main

import "log"

// 节点，包含两部分
type node struct {
	value int	//值
	next *node	//指向下一个节点的指针
}

type single_link_list struct {
	head *node
}

//创建新节点 这个没想到
func createNewNode(val int) *node {
	return &node{val, nil}
}

//添加节点
func (ll *single_link_list)addAtEnd(val int){
	// 首先创建一个新的节点
	n := createNewNode(val)

	// 因为这里搞了半天  报错地址错误
	if ll.head == nil {
		ll.head = n
		return
	}

	//找到最后一个节点
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = n
}

//遍历输出单链表的所有节点
func (ll *single_link_list)ergodic(){
	cur := ll.head

	for cur.next != nil {

		log.Println(cur.value)

		cur = cur.next
	}

	log.Println(cur.value)
}

func main() {
	// 初始化一个单链表
	ll := single_link_list{}

	ll.addAtEnd(1)
	ll.addAtEnd(2)
	ll.addAtEnd(3)
	ll.addAtEnd(4)
	ll.addAtEnd(5)
	ll.addAtEnd(6)

	ll.ergodic()
}
