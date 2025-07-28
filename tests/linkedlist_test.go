package tests

import (
	c "data-visualizer/structures"
	"fmt"
)

func LinkedListTest() { //链表测试函数，用于测试链表的功能是否正常
	l := c.LinkedList{}      //创建一个链表
	l.InsertHead(1)          //在链表头部插入元素1
	l.InsertTail(2)          //在链表尾部插入元素2
	l.InsertAt(1, 3)         //在索引1处插入元素3
	fmt.Println(l.Display()) //显示链表中的元素
}
