package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFiles(root string) ([]string, error) {
	var fps []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ディレクトリは除く
		if !info.IsDir() {
			fps = append(fps, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fps, nil
}

func LibTest() {
	fmt.Println("LibTest")
}
