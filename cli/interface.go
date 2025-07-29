package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ShowMainMenu 显示主菜单
func ShowMainMenu() {
	fmt.Println("数据结构可视化工具 v1.0")
	fmt.Println("======================")
	fmt.Println()
	fmt.Println("请选择数据结构:")
	fmt.Println("1. 数组 (Array)")
	fmt.Println("2. 链表 (LinkedList)")
	fmt.Println("3. 栈 (Stack)")
	fmt.Println("4. 队列 (Queue)")
	fmt.Println("5. 退出")
	fmt.Println()
	fmt.Print("请输入选择 (1-5): ")
}

// Run 启动CLI交互界面
func Run() {
	scanner := bufio.NewScanner(os.Stdin) // 创建一个新的Scanner，用于读取用户输入
	for {
		ShowMainMenu()

		// 读取用户输入
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text()) // 去除输入中的空格
		choice, err := strconv.Atoi(input)         // 将输入转换为整数
		if err != nil {
			fmt.Println("无效输入，请输入数字 1-5")
			continue
		}

		// 使用switch case处理选择
		switch choice {
		case 1:
			fmt.Println("您选择了数组 (Array)")
			RunArrayUI() // 调用数组UI
		case 2:
			fmt.Println("您选择了链表 (LinkedList)")
			RunLinkedListUI() // 调用链表UI
		// 在case 3和case 4中添加以下代码
		case 3:
			fmt.Println("您选择了栈 (Stack)")
			RunStackUI() // 调用栈UI
		case 4:
			fmt.Println("您选择了队列 (Queue)")
			RunQueueUI() // 调用队列UI

		case 5:
			fmt.Println("退出程序，再见！")
			return
		default:
			fmt.Println("无效选择，请输入 1-5 之间的数字")
		}

		fmt.Println("\n按回车键继续...")
		scanner.Scan()             // 等待用户按回车
		fmt.Print("\033[H\033[2J") // 清屏
	}
}
