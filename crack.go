package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// 读取要替换的文本和替换文本
	originalText, err := os.ReadFile("src.js")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	replacementText, err := os.ReadFile("dst.js")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	// 读取目标文件
	targetFileContent, err := os.ReadFile("/opt/Termius/resources/app/js/background-process.js")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	// 使用regexp进行替换
	// 注意：这里的正则表达式需要根据实际文本进行调整
	re := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(string(originalText)))
	replacedContent := re.ReplaceAll(targetFileContent, replacementText)

	// 将替换后的内容写回文件
	if err := os.WriteFile("/opt/Termius/resources/app/js/background-process.js", replacedContent, 0666); err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Println("替换完成。")
}
