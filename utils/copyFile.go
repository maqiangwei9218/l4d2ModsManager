package utils

import (
	"bufio"
	"io"
	"os"
)

func CopyFile(from string, to string) error {
	// 打开源文件
	source, err := os.Open(from)
	if err != nil {
		return err
	}
	defer source.Close()

	// 创建目标文件
	destination, err := os.Create(to)
	if err != nil {
		return err
	}
	defer destination.Close()

	// 使用带缓冲的读写
	bufReader := bufio.NewReader(source)
	bufWriter := bufio.NewWriter(destination)

	// 复制内容
	_, err = io.Copy(bufWriter, bufReader)
	if err != nil {
		return err
	}

	// 确保所有缓冲数据都写入
	return bufWriter.Flush()
}
