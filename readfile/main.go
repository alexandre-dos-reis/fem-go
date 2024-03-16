package main

import (
	fileutils "fem-go/readfile/file-utils"
	"fmt"
	"os"
)

func main() {
	rootDir, _ := os.Getwd()
	filePath := rootDir + "/data/text.txt"

	c, err := fileutils.ReadTextFile(filePath)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// `Sprintf` is used everywhere
	newContent := fmt.Sprintf("Original: %v\n Double original: %v %v", c, c, c)
	fileutils.WriteToFile(filePath+".output.txt", newContent)
	fmt.Println(c)
	fileutils.WriteToFile(filePath, newContent)
}
