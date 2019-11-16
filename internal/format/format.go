package format

import "fmt"

const (
	// Command color code
	Color_red     = uint8(iota + 91)
	Color_green   //	綠
	Color_yellow  //	黃
	Color_blue    // 	藍
	Color_magenta //	洋紅
)

// GetCMDColor 傳回在 terminal 上顯示顏色的字串
func GetCMDColor(color uint8, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}
