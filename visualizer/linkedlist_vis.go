package visualizer

import (
	"data-visualizer/structures"
	"strconv"
	"strings"
)

// VisualizeLinkedListInsertHead 可视化链表头部插入过程
func VisualizeLinkedListInsertHead(list *structures.LinkedList, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 在头部插入元素
	// result.WriteString(fmt.Sprintf("执行中: 在头部插入元素 %d\n", value))

	// 操作后状态
	// tempList := copyLinkedList(list)
	list.InsertHead(value)
	result.WriteString("操作后: ")
	result.WriteString(list.Display())

	return result.String()
}

// VisualizeLinkedListInsertTail 可视化链表尾部插入过程
func VisualizeLinkedListInsertTail(list *structures.LinkedList, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 在尾部插入元素
	// result.WriteString(fmt.Sprintf("执行中: 在尾部插入元素 %d\n", value))

	// 操作后状态
	// tempList := copyLinkedList(list)
	list.InsertTail(value)
	result.WriteString("操作后: ")
	result.WriteString(list.Display())

	return result.String()
}

// VisualizeLinkedListDeleteHead 可视化链表头部删除过程
func VisualizeLinkedListDeleteHead(list *structures.LinkedList) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 删除头部元素
	// result.WriteString("执行中: 删除头部元素\n")

	// 操作后状态
	// tempList := copyLinkedList(list)
	err := list.DeleteHead()
	if err != nil {
		result.WriteString("操作后: 无法删除 - " + err.Error())
	} else {
		result.WriteString("操作后: ")
		result.WriteString(list.Display())
	}

	return result.String()
}

// VisualizeLinkedListDeleteTail 可视化链表尾部删除过程
func VisualizeLinkedListDeleteTail(list *structures.LinkedList) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 执行中状态 - 删除尾部元素
	// result.WriteString("执行中: 删除尾部元素\n")

	// 操作后状态
	// tempList := copyLinkedList(list)
	err := list.DeleteTail()
	if err != nil {
		result.WriteString("操作后: 无法删除 - " + err.Error())
	} else {
		result.WriteString("操作后: ")
		result.WriteString(list.Display())
	}

	return result.String()
}

// VisualizeLinkedListInsert 可视化链表插入过程
func VisualizeLinkedListInsert(list *structures.LinkedList, index, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 插入位置指示
	// result.WriteString(createPositionIndicator(list.Size(), index))
	// result.WriteString("\n\n")

	// 执行中状态 - 在指定位置插入元素
	// result.WriteString(fmt.Sprintf("执行中: 在位置 %d 插入元素 %d\n", index, value))

	// 操作后状态
	// tempList := copyLinkedList(list)
	list.InsertAt(index, value)
	result.WriteString("操作后: ")
	result.WriteString(list.Display())

	return result.String()
}

// VisualizeLinkedListDelete 可视化链表删除过程
func VisualizeLinkedListDelete(list *structures.LinkedList, index int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(list.Display())
	result.WriteString("\n\n")

	// 删除位置指示
	// result.WriteString(createPositionIndicator(list.Size(), index))
	// result.WriteString("\n\n")

	// 执行中状态 - 删除指定位置元素
	// result.WriteString(fmt.Sprintf("执行中: 删除位置 %d 的元素\n", index))

	// 操作后状态
	// tempList := copyLinkedList(list)
	list.DeleteAt(index)
	result.WriteString("操作后: ")
	result.WriteString(list.Display())

	return result.String()
}

// 链表专用辅助函数
func copyLinkedList(list *structures.LinkedList) *structures.LinkedList {
	newList := &structures.LinkedList{}
	display := list.Display()
	parts := strings.Split(display, "->")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "head" || part == "NULL" {
			continue
		}

		if strings.Contains(part, "[") && strings.Contains(part, "]") {
			valStr := strings.Trim(part, "[] ")
			val, _ := strconv.Atoi(valStr)
			newList.InsertTail(val)
		}
	}

	return newList
}
