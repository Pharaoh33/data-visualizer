package structures

import "fmt"

type Node struct { //节点
	data int   //数据
	next *Node //下一个节点
}

type LinkedList struct {
	head *Node //头节点
	size int   //大小
}

func (l *LinkedList) InsertHead(value int) {
	node := &Node{ //创建一个新节点
		data: value,
		next: l.head, //将新节点的next指向头节点
	}
	l.head = node //将新节点作为头节点
	l.size++      //大小加1
}

func (l *LinkedList) InsertTail(value int) {
	node := &Node{
		data: value,
		next: nil,
	}
	if l.head == nil { //如果头节点为空，将新节点作为头节点
		l.head = node
		l.size++
		return
	}
	cur := l.head //遍历到尾节点
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node //将新节点作为尾节点
	l.size++
}

func (l *LinkedList) InsertAt(index, value int) error {
	//检查索引是否有效
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of range")
	}
	if index == 0 { //如果索引为0，插入到头节点
		l.InsertHead(value)
		return nil
	}
	if index == l.size { //如果索引为l.size，插入到尾节点
		l.InsertTail(value)
		return nil
	}
	node := &Node{ //创建一个新节点
		data: value,
		next: nil,
	}
	cur := l.head //遍历到索引为index的节点
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	node.next = cur.next //将新节点的next指向当前节点的next
	cur.next = node      //将当前节点的next指向新节点
	l.size++             //大小加1
	return nil
}

func (l *LinkedList) DeleteHead() error {
	if l.head == nil { //如果头节点为空，返回错误
		return fmt.Errorf("list is empty")
	}
	l.head = l.head.next //将头节点的next作为头节点
	l.size--             //大小减1
	return nil
}

func (l *LinkedList) DeleteTail() error {
	if l.head == nil { //如果头节点为空，返回错误
		return fmt.Errorf("list is empty")
	}
	if l.head.next == nil { //如果头节点的next为空，将头节点置为nil
		l.head = nil
		l.size--
		return nil
	}
	cur := l.head //遍历到尾节点
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil //将尾节点的next置为nil
	l.size--       //大小减1
	return nil
}

func (l *LinkedList) DeleteAt(index int) error {
	//检查索引是否有效
	if index < 0 || index >= l.size {
		return fmt.Errorf("index out of range")
	}
	if index == 0 { //如果索引为0，删除头节点
		return l.DeleteHead()
	}
	if index == l.size-1 { //如果索引为l.size-1，删除尾节点
		return l.DeleteTail()
	}
	cur := l.head //遍历到索引为index的节点
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next //将当前节点的next指向当前节点的next的next
	l.size--                 //大小减1
	return nil
}

func (l *LinkedList) Find(value int) int {
	cur := l.head //遍历到头节点
	for i := 0; i < l.size; i++ {
		if cur.data == value { //如果数据等于value，返回索引
			return i
		}
		cur = cur.next //将cur指向cur的next
	}
	return -1 //返回-1，表示没有找到
}

func (l *LinkedList) Size() int { //返回大小
	return l.size
}

func (l *LinkedList) Display() string { //返回字符串
	if l.head == nil {
		return "head -> NULL"
	}
	// 用于拼接结果字符串
	result := "head -> "
	current := l.head
	for current != nil {
		// 拼接当前节点数据，格式为 [data]
		result += fmt.Sprintf("[%d] -> ", current.data)
		current = current.next
	}
	// 替换最后一个 " -> " 为 " -> NULL"
	return result + "NULL"
}
