package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

type FileNode struct {
	path string
	size int64
}

func (fn FileNode) String() string {
	return fmt.Sprintf("%10s  %s", SizeString(fn.size), fn.path);
}

func SizeString(size int64) string {
	s := float64(size);
	r := fmt.Sprintf("%.0f B ", s)

	if s > 1024.0 {
		s /= 1024.0
		r = fmt.Sprintf("%.3f KB", s)
	}

	if s > 1024.0 {
		s /= 1024.0
		r = fmt.Sprintf("%.3f MB", s)
	}

	if s > 1024.0 {
		s /= 1024.0
		r = fmt.Sprintf("%.3f GB", s)
	}

	return r;
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if pwd[len(pwd) - 1] != os.PathSeparator {
		pwd = fmt.Sprintf("%s%c", pwd, os.PathSeparator)
	}

	totalSize := int64(0)
	filesList := []FileNode{}
	err = filepath.Walk(pwd, func(path string, f os.FileInfo, err error) error {
		fi, e := os.Stat(path);
		size := int64(0)

		if e != nil {
			size = 0
		} else {
			size = fi.Size()
		}

		totalSize += size
		path = fmt.Sprintf(".%c%s", os.PathSeparator, strings.TrimPrefix(path, pwd))
		filesList = append(filesList, FileNode{path, size})
		return nil
	})

	for _, file := range filesList {
		fmt.Println(file)
	}

	fmt.Println()
	fmt.Println(pwd)
	fmt.Printf("Files/Directories: %d    Total size: %s\n",
		len(filesList), SizeString(totalSize))
}
