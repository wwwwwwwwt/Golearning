/*
 * @Author: zzzzztw
 * @Date: 2023-05-27 16:06:40
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-05-27 16:37:55
 * @FilePath: /Golearning/example/basego/any.go
 */
package main

import "fmt"

// 定义一个可以包含任意类型值的单链表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &element[T]{nil, v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{nil, v}
		l.tail = l.tail.next
	}
}

func (l *List[T]) GetAll() []T {
	var ele []T

	for i := l.head; i != nil; i = i.next {
		ele = append(ele, i.val)
	}
	return ele
}

func main() {
	var li List[int]

	li.Push(10)
	li.Push(20)
	li.Push(30)

	num := li.GetAll()
	fmt.Println(num)
}
