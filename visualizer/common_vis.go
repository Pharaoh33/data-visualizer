package visualizer

import (
	"strings"
)

// createPositionIndicator 创建位置指示器（共用）
func createPositionIndicator(size, index int) string {
	var result strings.Builder
	for i := 0; i < index; i++ {
		result.WriteString("  ") // 每个位置占2个字符
	}
	result.WriteString("↑")
	return result.String()
}
