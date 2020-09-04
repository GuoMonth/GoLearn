package xml

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读取文件
func readFile(filePath string) string {
	fileObj, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err", filePath, err)
	}
	defer fileObj.Close()
	content, err := ioutil.ReadAll(fileObj)
	if err != nil {
		fmt.Println("read file err", filePath, err)
	}

	return string(content)
}
