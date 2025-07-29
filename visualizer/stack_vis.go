package visualizer

import (
	"data-visualizer/structures"
	"fmt"
	"strings"
)

// VisualizeStackPush 可视化栈推入过程
func VisualizeStackPush(stack *structures.Stack, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前:\n")
	result.WriteString(stack.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 推入新元素
	result.WriteString(fmt.Sprintf("执行中: 推入元素 %d\n", value))

	// 操作后状态
	// tempStack := copyStack(stack)
	stack.Push(value)
	result.WriteString("操作后:\n")
	result.WriteString(stack.Display())

	return result.String()
}

// VisualizeStackPop 可视化栈弹出过程
func VisualizeStackPop(stack *structures.Stack) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前:\n")
	result.WriteString(stack.Display())
	result.WriteString("\n\n")

	// 获取栈顶元素
	topValue, _ := stack.Peek()

	// 执行中状态 - 弹出元素
	result.WriteString(fmt.Sprintf("执行中: 弹出元素 %d\n", topValue))

	// 操作后状态
	// stack := copyStack(stack)
	_, err := stack.Pop()
	if err != nil {
		return err.Error()
	}
	result.WriteString("操作后:\n")
	result.WriteString(stack.Display())

	return result.String()
}

// 栈专用辅助函数
func copyStack(stack *structures.Stack) *structures.Stack {
	newStack := &structures.Stack{}
	temp := []int{}

	// 弹出所有元素并保存
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		temp = append(temp, val)
	}

	// 将元素推回原栈和新栈
	for i := len(temp) - 1; i >= 0; i-- {
		stack.Push(temp[i])
		newStack.Push(temp[i])
	}

	return newStack
}
