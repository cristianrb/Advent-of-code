package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Directory struct {
	Name           string
	Parent         *Directory
	Subdirectories map[string]*Directory
	Files          []File
	Size           int
}

type File struct {
	Name string
	Size int
}

func main() {
	input, _ := io.ReadLines("Day-07/input_01")
	rootFolder := createFileSystem(input)
	res1 := findDirectoriesWithSizeLessThan(rootFolder, 100000)

	minSize := 30000000 - (70000000 - rootFolder.Size)
	res2 := findSmallestDirectoryToDelete(rootFolder, minSize)
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
}

func createFileSystem(input []string) *Directory {
	root := &Directory{
		Name:           "/",
		Subdirectories: map[string]*Directory{},
	}
	actualDir := root
	for _, cmd := range input {
		if unicode.IsDigit(rune(cmd[0])) {
			fileSplit := strings.Split(cmd, " ")
			filename := fileSplit[1]
			filesize, _ := strconv.Atoi(fileSplit[0])
			file := File{
				Name: filename,
				Size: filesize,
			}
			actualDir.Files = append(actualDir.Files, file)
			actualDir.Size += filesize
		} else if strings.Contains(cmd, "dir") {
			dirName := strings.Split(cmd, "dir ")[1]
			dir := &Directory{
				Name:           dirName,
				Parent:         actualDir,
				Subdirectories: map[string]*Directory{},
			}
			actualDir.Subdirectories[dirName] = dir
		} else if strings.Contains(cmd, "cd") {
			newDir := strings.Split(cmd, "cd ")[1]
			if newDir == ".." {
				actualDir.Parent.Size += actualDir.Size
				actualDir = actualDir.Parent
			} else {
				actualDir = actualDir.Subdirectories[newDir]
			}
		}
	}
	actualDir.Parent.Size += actualDir.Size
	return root
}

func findDirectoriesWithSizeLessThan(root *Directory, n int) int {
	total := 0
	for _, dir := range root.Subdirectories {
		if dir.Subdirectories != nil {
			total += findDirectoriesWithSizeLessThan(dir, n)
		}
		if dir.Size < n {
			total += dir.Size
		}
	}

	return total
}

func findSmallestDirectoryToDelete(root *Directory, minSize int) int {
	smallest := 999999999999999
	for _, dir := range root.Subdirectories {
		if dir.Subdirectories != nil {
			smallestSub := findSmallestDirectoryToDelete(dir, minSize)
			if smallestSub < smallest {
				smallest = smallestSub
			}
		}

		if dir.Size >= minSize && dir.Size <= smallest {
			smallest = dir.Size
		}
	}

	return smallest
}
