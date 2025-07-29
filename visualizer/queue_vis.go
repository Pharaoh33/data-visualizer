package visualizer

import (
	"data-visualizer/structures"
	"fmt"
	"strings"
)

// VisualizeQueueEnqueue 可视化队列入队过程
func VisualizeQueueEnqueue(queue *structures.Queue, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(queue.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 入队新元素
	result.WriteString(fmt.Sprintf("执行中: 入队元素 %d\n", value))

	// 操作后状态
	tempQueue := copyQueue(queue)
	tempQueue.Enqueue(value)
	result.WriteString("操作后: ")
	result.WriteString(tempQueue.Display())

	return result.String()
}

// VisualizeQueueDequeue 可视化队列出队过程
func VisualizeQueueDequeue(queue *structures.Queue) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(queue.Display())
	result.WriteString("\n\n")

	// 获取队头元素
	frontValue, _ := queue.Front()

	// 执行中状态 - 出队元素
	result.WriteString(fmt.Sprintf("执行中: 出队元素 %d\n", frontValue))

	// 操作后状态
	tempQueue := copyQueue(queue)
	tempQueue.Dequeue()
	result.WriteString("操作后: ")
	result.WriteString(tempQueue.Display())

	return result.String()
}

// 队列专用辅助函数
func copyQueue(queue *structures.Queue) *structures.Queue {
	newQueue := &structures.Queue{}
	temp := []int{}

	// 出队所有元素并保存
	for !queue.IsEmpty() {
		val, _ := queue.Dequeue()
		temp = append(temp, val)
	}

	// 将元素入队回原队列和新队列
	for _, val := range temp {
		queue.Enqueue(val)
		newQueue.Enqueue(val)
	}

	return newQueue
}
