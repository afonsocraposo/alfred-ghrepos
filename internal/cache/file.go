package cache

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MkdirIfNotExist(path string) bool {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func readTextBytes(path string) ([]byte, bool) {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Println(err)
		return []byte{}, false
	}

	return content, true
}

func writeTextFile(bytes []byte, path string) bool {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()
	_, err = f.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
