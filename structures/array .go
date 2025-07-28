package structures

import (
	"fmt"
	"strings"
)

type Array struct {
	data     []int //存放数据
	capacity int   //容量
	size     int   //实际大小
}

func NewArray(capacity int) *Array { //创建数组
	return &Array{ //返回一个数组
		data:     make([]int, capacity), //创建一个容量为capacity的数组
		capacity: capacity,
		size:     0,
	}
}

func (a *Array) Insert(index, value int) error {
	if a.size == a.capacity {
		return fmt.Errorf("array is full") //返回数组已满
	}
	if index < 0 || index > a.size {
		return fmt.Errorf("index out of range") //返回索引超出范围
	}
	for i := a.size - 1; i >= index; i-- { //从后往前遍历
		a.data[i+1] = a.data[i] //将元素向后移动
	}
	a.data[index] = value //将元素插入到index位置
	a.size++              //实际大小加1
	return nil            //返回nil
}

func (a *Array) Delete(index int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range") //返回索引超出范围
	}
	for i := index; i < a.size-1; i++ { //从前往后遍历
		a.data[i] = a.data[i+1] //将元素向前移动
	}
	a.size--   //实际大小减1
	return nil //返回nil
}

func (a *Array) Get(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, fmt.Errorf("index out of range") //返回索引超出范围
	}
	return a.data[index], nil //返回元素
}

func (a *Array) Set(index, value int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range") //返回索引超出范围
	}
	a.data[index] = value //将元素设置为value
	return nil            //返回nil
}

func (a *Array) Find(value int) int {
	for i := 0; i < a.size; i++ { //从前往后遍历
		if a.data[i] == value { //如果元素等于value
			return i //返回索引
		}
	}
	return -1 //返回-1
}

func (a *Array) Display() string {
	var topBorder strings.Builder
	var dataLine strings.Builder
	var bottomBorder strings.Builder
	// var indexLine strings.Builder

	// 构建顶部边框
	topBorder.WriteString("┌")
	for i := 0; i < a.capacity; i++ {
		if i < a.capacity-1 {
			topBorder.WriteString("─────┬")
		} else {
			// 最后一个位置直接拼接 ┐
			topBorder.WriteString("─────┐")
		}
	}
	topBorder.WriteString("\n")

	// 构建数据行
	dataLine.WriteString("│")
	for i := 0; i < a.capacity; i++ {
		if i < a.size {
			dataLine.WriteString(fmt.Sprintf("  %d  │", a.data[i]))
		} else {
			dataLine.WriteString("     │")
		}
	}
	dataLine.WriteString("\n")

	// 构建底部边框
	bottomBorder.WriteString("└")
	for i := 0; i < a.capacity; i++ {
		if i < a.capacity-1 {
			bottomBorder.WriteString("─────┴")
		} else {
			// 最后一个位置直接拼接 ┘
			bottomBorder.WriteString("─────┘")
		}
	}
	bottomBorder.WriteString("\n")

	indexLine := "  "
	for i := 0; i < a.capacity; i++ {
		// 每个索引的起始位置 = 2（左侧空格） + i*6（每个单元格占6个字符）
		startPos := 2 + i*6
		// 填充空格到起始位置
		indexLine += strings.Repeat(" ", startPos-len(indexLine))
		// 添加索引
		indexLine += fmt.Sprintf("[%d]", i)
	}

	// // 构建索引行
	// for i := 0; i < a.capacity; i++ {
	// 	indexLine.WriteString(fmt.Sprintf("[%d]   ", i))
	// }

	// 合并所有部分
	return topBorder.String() + dataLine.String() + bottomBorder.String() + indexLine
}

func (a *Array) Size() int {
	return a.size //返回实际大小
}
