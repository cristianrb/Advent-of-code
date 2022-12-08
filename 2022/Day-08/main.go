package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
)

func main() {
	grid := io.ReadIntLinesAs2D("Day-08/input_01")
	res1, res2 := visibleTrees(grid)
	fmt.Printf("Part 1: %d\n", res1)
	fmt.Printf("Part 2: %d\n", res2)
}

func visibleTrees(grid [][]int) (int, int) {
	nVisibleTrees := 0
	highestScenicScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tree := grid[i][j]
			if isOutside(grid, i, j) {
				nVisibleTrees++
			} else {
				visibleL, lss := checkLeftAndScenicScore(tree, grid, i, j)
				visibleR, rss := checkRightAndScenicScore(tree, grid, i, j)
				visibleT, tss := checkTopAndScenicScore(tree, grid, i, j)
				visibleB, bss := checkBottomAndScenicScore(tree, grid, i, j)
				if visibleL || visibleR || visibleB || visibleT {
					nVisibleTrees++
				}
				score := lss * rss * bss * tss
				if score > highestScenicScore {
					highestScenicScore = score
				}
			}
		}
	}

	return nVisibleTrees, highestScenicScore
}

func isOutside(grid [][]int, i int, j int) bool {
	return i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1
}

func checkLeftAndScenicScore(tree int, grid [][]int, i int, j int) (bool, int) {
	sc := 0
	for k := j - 1; k >= 0; k-- {
		sc++
		if tree <= grid[i][k] {
			return false, sc
		}
	}
	return true, sc
}

func checkRightAndScenicScore(tree int, grid [][]int, i int, j int) (bool, int) {
	sc := 0
	for k := j + 1; k < len(grid[0]); k++ {
		sc++
		if tree <= grid[i][k] {
			return false, sc
		}
	}
	return true, sc
}

func checkTopAndScenicScore(tree int, grid [][]int, i int, j int) (bool, int) {
	sc := 0
	for k := i - 1; k >= 0; k-- {
		sc++
		if tree <= grid[k][j] {
			return false, sc
		}
	}
	return true, sc
}

func checkBottomAndScenicScore(tree int, grid [][]int, i int, j int) (bool, int) {
	sc := 0
	for k := i + 1; k < len(grid); k++ {
		sc++
		if tree <= grid[k][j] {
			return false, sc
		}
	}
	return true, sc
}
