package stringOps

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

/**
 *	字符串操作
 *
 */

//StrOpsRun ...开始运行
func StrOpsRun() {
	strForRange()
}

// 使用 for range 遍历字符串
func strForRange() {
	str := "好好学习，天天向上！Good"
	ruStr := []rune(str) // 统一转换为utf-8编码
	for i, c := range ruStr {
		fmt.Println(i, c, ruStr[i], c == ruStr[i], string(c), string(ruStr[i]))
	}

}

func strSlicForRange() {
	s := []byte("UP 天 DAY!")
	var buffer bytes.Buffer
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		buffer.WriteString(fmt.Sprintf("r:%v,width:%v\n", r, width))
		i += width
	}
	fmt.Printf(buffer.String())

	si := []int{1, 2, 3, 4, 5}
	si1 := si[0:]

	for _, s := range si1 {
		fmt.Println("s", s)
	}

}
