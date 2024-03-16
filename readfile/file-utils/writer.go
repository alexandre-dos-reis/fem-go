package fileutils

import "os"

func WriteToFile(filename string, content string) error {
	// https://pkg.go.dev/os@go1.20.5#WriteFile
	return os.WriteFile(filename, []byte(content), 0644)
}
