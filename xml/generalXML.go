package xml

import (
	"fmt"
	"os"
)

func writeLogFile(filePath string, str string) {
	var f *os.File
	var err error
	if !exitsFile(filePath) {
		f, err = os.Open(filePath)
	} else {
		f, err = os.Create(filePath)
	}

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	f.WriteString(str)
}

func exitsFile(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}
