package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"data-visualizer/structures"
)

// ShowLinkedListMenu 显示链表操作菜单
func ShowLinkedListMenu() {
	fmt.Println("=== 链表操作 ===")
	fmt.Println("可用操作:")
	fmt.Println("1. 头部插入 (insert head)")
	fmt.Println("2. 尾部插入 (insert tail)")
	fmt.Println("3. 指定位置插入 (insert at)")
	fmt.Println("4. 头部删除 (delete head)")
	fmt.Println("5. 尾部删除 (delete tail)")
	fmt.Println("6. 指定位置删除 (delete at)")
	fmt.Println("7. 查找 (find)")
	fmt.Println("8. 显示 (show)")
	fmt.Println("9. 返回主菜单 (back)")
	fmt.Println()
	fmt.Print("请输入操作: ")
}

// RunLinkedListUI 启动链表操作界面
func RunLinkedListUI() {
	scanner := bufio.NewScanner(os.Stdin)
	list := structures.LinkedList{} // 初始化链表
	
	for {
		// 显示当前链表状态
		fmt.Println("当前链表:", list.Display())
		fmt.Println()
		
		ShowLinkedListMenu()
		
		// 读取用户输入
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效输入，请输入数字 1-9")
			continue
		}
		
		// 使用switch处理链表操作选择
		switch choice {
		case 1:
			// 头部插入
			fmt.Print("请输入插入值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			list.InsertHead(value)
			fmt.Printf("已在头部插入值 %d\n", value)
		case 2:
			// 尾部插入
			fmt.Print("请输入插入值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			list.InsertTail(value)
			fmt.Printf("已在尾部插入值 %d\n", value)
		case 3:
			// 指定位置插入
			fmt.Print("请输入插入位置: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Print("请输入插入值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			err := list.InsertAt(index, value)
			if err != nil {
				fmt.Println("插入失败:", err)
			} else {
				fmt.Printf("已在位置 %d 插入值 %d\n", index, value)
			}
		case 4:
			// 头部删除
			err := list.DeleteHead()
			if err != nil {
				fmt.Println("删除失败:", err)
			} else {
				fmt.Println("已删除头部元素")
			}
		case 5:
			// 尾部删除
			err := list.DeleteTail()
			if err != nil {
				fmt.Println("删除失败:", err)
			} else {
				fmt.Println("已删除尾部元素")
			}
		case 6:
			// 指定位置删除
			fmt.Print("请输入删除位置: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			err := list.DeleteAt(index)
			if err != nil {
				fmt.Println("删除失败:", err)
			} else {
				fmt.Printf("已删除位置 %d 的元素\n", index)
			}
		case 7:
			// 查找操作
			fmt.Print("请输入查找值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			index := list.Find(value)
			if index != -1 {
				fmt.Printf("值 %d 找到，位置为 %d\n", value, index)
			} else {
				fmt.Printf("值 %d 未找到\n", value)
			}
		case 8:
			// 显示操作
			fmt.Println("链表元素:")
			fmt.Println(list.Display())
		case 9:
			// 返回主菜单
			fmt.Println("返回主菜单...")
			return
		default:
			fmt.Println("无效选择，请输入 1-9 之间的数字")
		}
		
		fmt.Println("\n按回车键继续...")
		scanner.Scan()
		fmt.Print("\033[H\033[2J") // 清屏
	}
}
