package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func panicTodo(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// 使用正则表达式匹配以 func 开始的函数，替换函数体为 panic("TODO")
	re := regexp.MustCompile(`(\n//[\s\S]*?func.*?{)[\s\S]*?\n}`)
	newContent := re.ReplaceAllString(string(content), "$1\n\t panic(\"TODO???\")\n}")
	fmt.Println(newContent)
	// 写入替换后的内容到文件
	err = os.WriteFile(file, []byte(newContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// mm
func processGoFiles(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go") {
			if info.Name() == "clear.go" {
				return nil
			}
			fmt.Printf("Processing: %s\n", path)
			err := panicTodo(path)
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			}
		}
		return nil
	})

	return err
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	err = processGoFiles(currentDir)
	if err != nil {
		fmt.Printf("Error processing Go files: %v\n", err)
		return
	}

	fmt.Println("All Go files processed successfully.")
}
