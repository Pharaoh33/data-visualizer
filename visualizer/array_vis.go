package visualizer

import (
	"data-visualizer/structures"
	"fmt"
	"strings"
)

// VisualizeArrayInsert 可视化数组插入过程
func VisualizeArrayInsert(arr *structures.Array, index, value int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(arrayToSimpleString(arr))
	result.WriteString("\n")

	// 插入位置指示
	result.WriteString(createPositionIndicator(arr.Size(), index))
	result.WriteString("\n\n")

	// 执行中状态 - 为新元素腾出空间
	result.WriteString("执行中: ")
	result.WriteString(arrayInsertStepString(arr, index))
	result.WriteString("\n")
	result.WriteString(createPositionIndicator(arr.Size(), index))
	result.WriteString("\n\n")

	// 操作后状态
	tempArr := copyArray(arr)
	tempArr.Insert(index, value)
	result.WriteString("操作后: ")
	result.WriteString(arrayToSimpleString(tempArr))

	return result.String()
}

// VisualizeArrayDelete 可视化数组删除过程
func VisualizeArrayDelete(arr *structures.Array, index int) string {
	var result strings.Builder

	// 操作前状态
	result.WriteString("操作前: ")
	result.WriteString(arrayToSimpleString(arr))
	result.WriteString("\n")

	// 删除位置指示
	result.WriteString(createPositionIndicator(arr.Size(), index))
	result.WriteString("\n\n")

	// 获取要删除的值
	deletedValue, _ := arr.Get(index)

	// 执行中状态 - 标记要删除的元素
	result.WriteString("执行中: ")
	result.WriteString(arrayDeleteStepString(arr, index))
	result.WriteString("\n")
	result.WriteString(createPositionIndicator(arr.Size(), index))
	result.WriteString("\n\n")

	// 操作后状态
	tempArr := copyArray(arr)
	tempArr.Delete(index)
	result.WriteString("操作后: ")
	result.WriteString(arrayToSimpleString(tempArr))
	result.WriteString(fmt.Sprintf("\n已删除元素: %d", deletedValue))

	return result.String()
}

// 数组专用辅助函数
func arrayToSimpleString(arr *structures.Array) string {
	var result strings.Builder
	for i := 0; i < arr.Size(); i++ {
		if i < arr.Size() {
			val, _ := arr.Get(i)
			result.WriteString(fmt.Sprintf("[%d]", val))
		} else {
			result.WriteString("[]")
		}
	}
	return result.String()
}

func arrayInsertStepString(arr *structures.Array, index int) string {
	tempArr := copyArray(arr)
	for i := tempArr.Size() - 1; i >= index; i-- {
		val, _ := tempArr.Get(i)
		tempArr.Delete(i)
		tempArr.Insert(i+1, val)
	}
	tempArr.Insert(index, 0)
	return arrayToSimpleString(tempArr)
}

func arrayDeleteStepString(arr *structures.Array, index int) string {
	tempArr := copyArray(arr)
	tempArr.Set(index, 0)
	return arrayToSimpleString(tempArr)
}

func copyArray(arr *structures.Array) *structures.Array {
	newArr := structures.NewArray(arr.Size())
	for i := 0; i < arr.Size(); i++ {
		val, _ := arr.Get(i)
		newArr.Insert(i, val)
	}
	return newArr
}
