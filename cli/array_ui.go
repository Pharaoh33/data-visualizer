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

// ShowArrayMenu 显示数组操作菜单
func ShowArrayMenu() {
	fmt.Println("数组操作菜单")
	fmt.Println("======================")
	fmt.Println("1. 插入 (insert)")
	fmt.Println("2. 删除 (delete)")
	fmt.Println("3. 查找 (find)")
	fmt.Println("4. 获取 (get)")
	fmt.Println("5. 设置 (set)")
	fmt.Println("6. 显示 (show)")
	fmt.Println("7. 返回主菜单 (back)")
	fmt.Println()
	fmt.Print("请输入操作: ")
}

// RunArrayUI 启动数组操作界面
func RunArrayUI() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := structures.NewArray(5) // 假设存在创建数组的函数

	for {

		// 显示当前数组状态
		fmt.Println("当前数组:")
		fmt.Println(arr.Display())

		ShowArrayMenu()

		// 读取用户输入
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效输入，请输入数字 1-7")
			continue
		}

		// 使用switch处理数组操作选择
		switch choice {
		case 1:
			// 插入操作
			fmt.Print("请输入插入索引: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Print("请输入插入值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeArrayInsert(arr, index, value))
			// err = arr.Insert(index, value)
			// if err != nil {
			// 	fmt.Printf("插入失败: %v\n", err)
			// 	continue
			// }
			// fmt.Printf("已在索引 %d 插入值 %d\n", index, value)
		case 2:
			// 删除操作
			fmt.Print("请输入删除索引: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Println("\n可视化过程:")
			fmt.Println(visualizer.VisualizeArrayDelete(arr, index))
			// err := arr.Delete(index)
			// if err != nil {
			// 	fmt.Printf("删除失败: %v\n", err)
			// 	continue
			// }
			// fmt.Printf("已删除索引 %d 的元素\n", index)
		case 3:
			// 查找操作
			fmt.Print("请输入查找值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			index := arr.Find(value)
			if index != -1 {
				fmt.Printf("值 %d 找到，索引为 %d\n", value, index)
			} else {
				fmt.Printf("值 %d 未找到\n", value)
			}
		case 4:
			// 获取操作
			fmt.Print("请输入获取索引: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			value, err := arr.Get(index)
			if err != nil {
				fmt.Printf("获取失败: %v\n", err)
				continue
			}
			fmt.Printf("索引 %d 的值为 %d\n", index, value)
		case 5:
			// 设置操作
			fmt.Print("请输入设置索引: ")
			scanner.Scan()
			index, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Print("请输入设置值: ")
			scanner.Scan()
			value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			err := arr.Set(index, value)
			if err != nil {
				fmt.Printf("设置失败: %v\n", err)
				continue
			}
			fmt.Printf("已将索引 %d 设置为 %d\n", index, value)
		case 6:
			// 显示操作
			fmt.Println("数组元素:")
			fmt.Println(arr.Display())
		case 7:
			// 返回主菜单
			fmt.Println("返回主菜单...")
			return
		default:
			fmt.Println("无效选择，请输入 1-7 之间的数字")
		}

		fmt.Println("\n按回车键继续...")
		scanner.Scan()
		fmt.Print("\033[H\033[2J") // 清屏
	}
}
