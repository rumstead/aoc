package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	cwd := newRootFS()
	// build the filesystem
	i := 0
	data := strings.Split(input, "\n")
	for i < len(data) {
		fmt.Printf("%d %s\n", i, data[i])
		if isCmd(data[i]) {
			if strings.Contains(data[i], "cd") {
				words := strings.Split(data[i], "cd")
				cwd = cd(cwd, words[1])
				i++
			} else if strings.Contains(data[i], "ls") {
				// parse the command output
				i++
				current := i
				for i < len(data) && !isCmd(data[i]) {
					// grab all the output lines
					i++
				}
				// back up one to process the command
				parseLS(data[current:i], cwd)
			}
		}
	}
	cwd = cd(cwd, "/")
	fmt.Println(totalDir(cwd))
}

func totalDir(node *file) int {
	size := 0
	queue := []*file{node}
	for i := 0; i < len(queue); i++ {
		child := queue[i]
		if !child.isDir && child.size <= 100000 {
			size += child.size
		} else if child.isDir {
			queue = append(queue, child.children...)
		}
	}
	return size
}

func parseLS(cmdOutput []string, cwd *file) {
	var fs []*file
	for i, line := range cmdOutput {
		words := strings.Fields(cmdOutput[i])
		if unicode.IsDigit(rune(line[0])) {
			// handle file
			fileSize, _ := strconv.Atoi(words[0])
			fs = append(fs, &file{
				name:     words[1],
				children: nil,
				parent:   nil,
				size:     fileSize,
				isDir:    false,
			})
			cwd.size += fileSize
			for node := cwd.parent; node != nil; {
				node.size += fileSize
				node = node.parent
			}
		} else {
			// directory
			fs = append(fs, &file{
				name:     words[1],
				children: nil,
				parent:   cwd,
				size:     0,
				isDir:    true,
			})
		}
	}
	cwd.children = append(cwd.children, fs...)
}

func cd(cwd *file, s string) *file {
	s = strings.TrimSpace(s)
	if s == "/" {
		for cwd.parent != nil {
			cwd = cwd.parent
		}
		return cwd
	} else if s == ".." {
		return cwd.parent
	} else {
		fmt.Printf("%s %v\n", s, cwd)
		for i, subfolder := range cwd.children {
			if subfolder.name == s && subfolder.isDir {
				return cwd.children[i]
			}
		}
	}
	return cwd
}

func isCmd(line string) bool {
	return strings.HasPrefix(line, "$")
}

type file struct {
	name     string
	children []*file
	parent   *file
	size     int
	isDir    bool
}

func newRootFS() *file {
	return &file{
		children: []*file{},
		parent:   nil,
		size:     0,
		name:     "/",
		isDir:    true,
	}
}
