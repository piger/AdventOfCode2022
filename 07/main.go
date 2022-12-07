package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Name     string
	Files    map[string]int
	Parent   *Dir
	Children map[string]*Dir
	Size     int
}

func NewDir(name string) *Dir {
	dir := Dir{
		Name:     name,
		Files:    make(map[string]int),
		Children: make(map[string]*Dir),
	}
	return &dir
}

func (d *Dir) AddSubdir(name string) *Dir {
	subdir, ok := d.Children[name]
	if !ok {
		subdir := NewDir(name)
		subdir.Parent = d
		d.Children[name] = subdir
	}
	return subdir
}

func run() error {
	fh, err := os.Open("sample")
	if err != nil {
		return err
	}
	defer fh.Close()

	root := NewDir("/")
	var cwd *Dir = root

	var lsMode bool

	s := bufio.NewScanner(fh)
	for s.Scan() {
		cmd := s.Text()
		if strings.HasPrefix(cmd, "$ cd") {
			lsMode = false
			fields := strings.Split(cmd, " ")
			dirName := fields[len(fields)-1]

			if dirName == ".." {
				// XXX check if Parent is not nil
				cwd = cwd.Parent
			} else if dirName == cwd.Name {
				continue
			} else {
				subdir := cwd.AddSubdir(dirName)
				cwd = subdir
			}
		} else if strings.HasPrefix(cmd, "$ ls") {
			lsMode = true
		} else if lsMode {
			fields := strings.Split(cmd, " ")
			if fields[0] == "dir" {
				cwd.AddSubdir(fields[1])
			} else {
				size, err := strconv.Atoi(fields[0])
				if err != nil {
					panic(err)
				}

				if _, ok := cwd.Files[fields[1]]; !ok {
					cwd.Files[fields[1]] = size
					cwd.Size += size
				}
			}
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Println("/:")
	printDir(root, 2)

	return nil
}

func printDir(dir *Dir, indent int) {
	prefix := strings.Repeat(" ", indent)

	for name, subdir := range dir.Children {
		fmt.Printf("%s%s:\n", prefix, name)
		printDir(subdir, indent+2)
	}

	for name, size := range dir.Files {
		fmt.Printf("%s%s: %d\n", prefix, name, size)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
