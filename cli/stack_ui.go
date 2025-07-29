package cli

import (
	"bufio"
	"data-visualizer/structures"
	"data-visualizer/visualizer"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ShowStackMenu 显示栈操作菜单
func ShowStackMenu() {
	fmt.Println("=== 栈操作 ===")
	fmt.Println("可用操作:")
	fmt.Println("1. 入栈 (push)")
	fmt.Println("2. 出栈 (pop)")
	fmt.Println("3. 查看栈顶 (peek)")
	fmt.Println("4. 显示栈 (show)")
	fmt.Println("5. 返回主菜单 (back)")
	fmt.Println()
	fmt.Print("请输入操作: ")
}

// RunStackUI 启动栈操作界面
func RunStackUI() {
	scanner := bufio.NewScanner(os.Stdin)
	stack := structures.Stack{} // 初始化栈

	for {
		// 显示当前栈状态
		fmt.Println("当前栈:")
		fmt.Println(stack.Display())
		fmt.Println()

		ShowStackMenu()

		// 读取用户输入
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效输入，请输入数字 1-5")
			continue
		}

		// 使用switch处理栈操作选择
		switch choice {
		// 入栈操作
		case 1:
			fmt.Print("请输入入栈值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeStackPush(&stack, value))
			// stack.Push(v,,%d\n", value)
		// 出栈操作
		case 2:
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeStackPop(&stack))
			// value, err := stack.Pop()
			// if err != nil {
			// 	fmt.Println("出栈失败:", err)
			// } else {
			// 	fmt.Printf("已出栈值 %d\n", value)
			// }
		case 3:
			// 查看栈顶
			value, err := stack.Peek()
			if err != nil {
				fmt.Println("栈为空:", err)
			} else {
				fmt.Printf("栈顶元素为 %d\n", value)
			}
		case 4:
			// 显示栈
			fmt.Println("栈元素:")
			fmt.Println(stack.Display())
		case 5:
			// 返回主菜单
			fmt.Println("返回主菜单...")
			return
		default:
			fmt.Println("无效选择，请输入 1-5 之间的数字")
		}

		fmt.Println("\n按回车键继续...")
		scanner.Scan()
		fmt.Print("\033[H\033[2J") // 清屏
	}
}
