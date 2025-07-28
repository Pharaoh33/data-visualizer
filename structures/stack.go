package structures

import "fmt"

type Stack struct { //栈
	data []int //存放数据
	size int   //实际大小
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value) //将元素添加到栈顶
	s.size++                       //实际大小加1
}

func (s *Stack) Pop() (int, error) { //弹出栈顶元素
	if s.size == 0 { //如果栈为空，返回错误
		return 0, fmt.Errorf("stack is empty") //返回错误
	}
	value := s.data[s.size-1]  //获取栈顶元素
	s.data = s.data[:s.size-1] //将栈顶元素弹出
	s.size--                   //实际大小减1
	return value, nil          //返回栈顶元素和nil
}

func (s *Stack) Peek() (int, error) { //获取栈顶元素
	if s.size == 0 { //如果栈为空，返回错误
		return 0, fmt.Errorf("stack is empty") //返回错误
	}
	return s.data[s.size-1], nil //返回栈顶元素和nil
}

func (s *Stack) IsEmpty() bool { //判断栈是否为空
	return s.size == 0 //返回栈是否为空
}

func (s *Stack) Size() int { //获取栈的大小
	return s.size //返回栈的大小
}

func (s *Stack) Display() string { //显示栈中的元素
	if s.size == 0 { //如果栈为空，返回空字符串
		return "栈为空" //返回空字符串
	}
	result := ""
	// 从栈顶往栈底遍历（因为栈顶是 top，数据存在 data 切片中，栈底是索引 0 方向 ）
	for i := s.size - 1; i >= 0; i-- {
		result += "| "
		if i < len(s.data) {
			result += fmt.Sprintf("%d", s.data[i])
		} else {
			result += " "
		}
		result += " |"
		if i == s.size-1 {
			result += " <- top"
		} else if i == 0 {
			result += " <- 栈底"
		}
		result += "\n"
	}
	return result
}
