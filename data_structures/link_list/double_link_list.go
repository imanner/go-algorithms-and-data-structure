package main

import "fmt"

type dnode struct {
	value int
	pre *dnode
	next *dnode
}

type double_link_list struct {
	head *dnode
}

func createNewDnode(val int) *dnode {
	return &dnode{val, nil, nil}
}

//添加节点
func (ll *double_link_list) addAtEnd(val int) {
	// 首先创建一个新的节点
	n := createNewDnode(val)

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

func (ll *double_link_list) delAtEnd() int {
	res := -1

	if ll.head == nil {
		return res
	}

	if ll.head.next == nil {
		res = ll.head.value
		ll.head.value = 0
		return res
	}

	cur := ll.head

	for ; cur.next.next != nil; cur = cur.next {
	}

	res = cur.next.value
	cur.next = nil
	return res
}



func main() {
	fmt.Println(123)
}