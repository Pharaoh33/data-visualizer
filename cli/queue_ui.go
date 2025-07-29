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

// ShowQueueMenu 显示队列操作菜单
func ShowQueueMenu() {
	fmt.Println("=== 队列操作 ===")
	fmt.Println("可用操作:")
	fmt.Println("1. 入队 (enqueue)")
	fmt.Println("2. 出队 (dequeue)")
	fmt.Println("3. 查看队头 (front)")
	fmt.Println("4. 显示队列 (show)")
	fmt.Println("5. 返回主菜单 (back)")
	fmt.Println()
	fmt.Print("请输入操作: ")
}

// RunQueueUI 启动队列操作界面
func RunQueueUI() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := structures.Queue{} // 初始化队列

	for {
		// 显示当前队列状态
		fmt.Println("当前队列:", queue.Display())
		fmt.Println()

		ShowQueueMenu()

		// 读取用户输入
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效输入，请输入数字 1-5")
			continue
		}

		// 使用switch处理队列操作选择
		switch choice {
		// 入队操作
		case 1:
			fmt.Print("请输入入队值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeQueueEnqueue(&queue, value))
			queue.Enqueue(value)
			fmt.Printf("已入队值 %d\n", value)
		// 出队操作
		case 2:
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeQueueDequeue(&queue))
			value, err := queue.Dequeue()
			if err != nil {
				fmt.Println("出队失败:", err)
			} else {
				fmt.Printf("已出队值 %d\n", value)
			}
		case 3:
			// 查看队头
			value, err := queue.Front()
			if err != nil {
				fmt.Println("队列为空:", err)
			} else {
				fmt.Printf("队头元素为 %d\n", value)
			}
		case 4:
			// 显示队列
			fmt.Println("队列元素:")
			fmt.Println(queue.Display())
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
