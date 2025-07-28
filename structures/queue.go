package structures

import (
	"fmt"
	"strings"
)

type Queue struct {
	data  []int
	front int //队头指针
	rear  int //队尾指针
	size  int
}

func (q *Queue) Enqueue(value int) { //入队
	q.data = append(q.data, value) //将元素添加到队尾
	if q.size != 0 {
		q.rear++
	}
	q.size++ //实际大小加1
}

func (q *Queue) Dequeue() (int, error) { //出队
	if q.size == 0 { //如果队列为空，返回错误
		return 0, fmt.Errorf("queue is empty") //返回错误
	}
	value := q.data[q.front] //获取队头元素
	q.data = q.data[1:]      //将队头元素弹出
	if q.size != 0 {
		q.rear--
	}
	q.size--          //实际大小减1
	return value, nil //返回队头元素和 nil
}

func (q *Queue) Front() (int, error) {
	if q.size == 0 { //如果队列为空，返回错误
		return 0, fmt.Errorf("queue is empty") //返回错误
	}
	return q.data[q.front], nil //返回队头元素和 nil
}

func (q *Queue) IsEmpty() bool { //判断队列是否为空
	return q.size == 0 //返回队列是否为空
}

func (q *Queue) Size() int { //获取队列的大小
	return q.size //返回队列的大小
}

func (q *Queue) Display() string { //显示队列中的元素
	if q.IsEmpty() {
		return "队列为空"
	}

	// 构建队列元素字符串
	elements := make([]string, 0, q.Size())
	for i := q.front; i <= q.rear; i++ {
		elements = append(elements, fmt.Sprintf("%d", q.data[i]))
	}
	elementStr := strings.Join(elements, " | ")

	// 构建可视化字符串
	return fmt.Sprintf(
		"出队 <- [%s] <- 入队\nfront: %d, rear: %d",
		elementStr, q.front, q.rear,
	)
}
