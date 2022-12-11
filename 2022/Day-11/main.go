package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"strconv"
	"strings"
)

type Operation int

const (
	Sum Operation = iota
	Multiply
)

type Monkey struct {
	Id              int
	items           []int
	Op              Operation
	OpBy            int
	TestDivisibleBy int
	TrueTo          int
	FalseTo         int
	Inspections     int
}

func main() {
	input, _ := io.ReadLines("Day-11/input_01")
	println(input)

	monkeys, _ := parseMonkeys(input)
	res1 := doRounds(monkeys, 20, func(n int) int {
		return n / 3
	})
	fmt.Printf("Part 1: %d\n", res1)

	monkeys2, lcm := parseMonkeys(input)
	res2 := doRounds(monkeys2, 10000, func(n int) int {
		return n % lcm
	})
	fmt.Printf("Part 2: %d\n", res2)
}

func doRounds(monkeys map[int]*Monkey, n int, op func(n int) int) int {
	for n > 0 {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			items := make([]int, len(monkey.items))
			copy(items, monkey.items)
			for i, _ := range items {
				monkey.Inspections++
				worryLevel := items[i]
				switch monkey.Op {
				case Sum:
					if monkey.OpBy == -1 {
						worryLevel += worryLevel
					} else {
						worryLevel += monkey.OpBy
					}
				case Multiply:
					if monkey.OpBy == -1 {
						worryLevel *= worryLevel
					} else {
						worryLevel *= monkey.OpBy
					}
				}
				worryLevel = op(worryLevel)
				if worryLevel%monkey.TestDivisibleBy == 0 {
					monkeys[monkey.TrueTo].items = append(monkeys[monkey.TrueTo].items, worryLevel)
				} else {
					monkeys[monkey.FalseTo].items = append(monkeys[monkey.FalseTo].items, worryLevel)
				}
				if len(monkey.items) == 1 {
					monkey.items = []int{}
				} else {
					monkey.items = append(monkey.items[:0], monkey.items[1:]...)
				}
			}
		}
		n--
	}
	return twoMostInspections(monkeys)
}

func twoMostInspections(monkeys map[int]*Monkey) int {
	max1, max2 := 0, 0
	for _, m := range monkeys {
		if max1 < m.Inspections {
			max2 = max1
			max1 = m.Inspections
		} else if max2 < m.Inspections {
			max2 = m.Inspections
		}
	}

	return max1 * max2
}

func parseMonkeys(input []string) (map[int]*Monkey, int) {
	monkeys := map[int]*Monkey{}
	lcm := 1
	var monkey *Monkey
	for idx, line := range input {
		if idx%7 == 0 {
			lineId := strings.Split(line, "Monkey ")[1]
			lineIdTrim := strings.TrimSuffix(lineId, ":")
			id, _ := strconv.Atoi(lineIdTrim)
			monkey = new(Monkey)
			monkey.Id = id
			monkeys[id] = monkey
		} else if idx%7 == 1 {
			items := strings.Split(strings.Split(line, "Starting items: ")[1], " ")
			itemsInt := []int{}
			for _, item := range items {
				itemInt, _ := strconv.Atoi(strings.TrimSuffix(item, ","))
				itemsInt = append(itemsInt, itemInt)
			}
			monkey.items = itemsInt
		} else if idx%7 == 2 {
			op := strings.Split(line, "new = ")[1]
			opBy := ""
			if strings.Contains(op, "*") {
				monkey.Op = Multiply
				opBy = strings.Split(op, "* ")[1]
			} else if strings.Contains(op, "+") {
				monkey.Op = Sum
				opBy = strings.Split(op, "+ ")[1]
			}
			if opBy == "old" {
				monkey.OpBy = -1 // -1 means old * old
			} else {
				opByNum, _ := strconv.Atoi(opBy)
				monkey.OpBy = opByNum
			}
		} else if idx%7 == 3 {
			test, _ := strconv.Atoi(strings.Split(line, "Test: divisible by ")[1])
			monkey.TestDivisibleBy = test
			lcm *= test
		} else if idx%7 == 4 {
			trueBranch, _ := strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
			monkey.TrueTo = trueBranch
		} else if idx%7 == 5 {
			falseBranch, _ := strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
			monkey.FalseTo = falseBranch
		}
	}

	return monkeys, lcm
}
