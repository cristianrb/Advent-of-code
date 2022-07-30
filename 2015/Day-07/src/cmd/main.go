package main

import (
	"cristianrb2015/io"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, _ := io.ReadLines("Day-07/inputs/input_01")
	wires := runCircuit(lines)
	println(wires["a"])
}

// Could be improved a lot.
func runCircuit(input []string) map[string]uint16 {
	wires := map[string]uint16{}

	for len(input) > 0 {
		idxToDel := []int{}
		for i, line := range input {
			if strings.Contains(line, "AND") {
				op1, op2, out, valid := andArguments(line, wires)
				if valid {
					wires[out] = op1 & op2
					idxToDel = append(idxToDel, i)
				}

			} else if strings.Contains(line, "OR") {
				arguments := strings.Split(line, " OR ")
				first := arguments[0]
				second := strings.Split(arguments[1], " -> ")[0]
				result := strings.Split(arguments[1], " -> ")[1]
				_, fExist := wires[first]
				_, sExist := wires[second]
				if fExist && sExist {
					wires[result] = wires[first] | wires[second]
					idxToDel = append(idxToDel, i)
				}

			} else if strings.Contains(line, "LSHIFT") {
				arguments := strings.Split(line, " LSHIFT ")
				first := arguments[0]
				_, fExist := wires[first]
				if fExist {
					second, _ := strconv.Atoi(strings.Split(arguments[1], " -> ")[0])
					result := strings.Split(arguments[1], " -> ")[1]
					wires[result] = wires[first] << second
					idxToDel = append(idxToDel, i)
				}

			} else if strings.Contains(line, "RSHIFT") {
				arguments := strings.Split(line, " RSHIFT ")
				first := arguments[0]
				_, fExist := wires[first]
				if fExist {
					second, _ := strconv.Atoi(strings.Split(arguments[1], " -> ")[0])
					result := strings.Split(arguments[1], " -> ")[1]
					wires[result] = wires[first] >> second
					idxToDel = append(idxToDel, i)
				}

			} else if strings.Contains(line, "NOT") {
				arguments := strings.Split(line, "NOT ")
				first := strings.Split(arguments[1], " -> ")[0]
				_, fExist := wires[first]
				if fExist {
					result := strings.Split(arguments[1], " -> ")[1]
					wires[result] = ^wires[first]
					idxToDel = append(idxToDel, i)
				}

			} else {
				arguments := strings.Split(line, " -> ")
				number, err := strconv.Atoi(arguments[0])
				if err != nil {
					_, fExist := wires[arguments[0]]
					if fExist {
						wires[arguments[1]] = wires[arguments[0]]
						idxToDel = append(idxToDel, i)
					}
				} else {
					wires[arguments[1]] = uint16(number)
					idxToDel = append(idxToDel, i)
				}

			}
		}
		sort.Slice(idxToDel, func(i, j int) bool {
			return idxToDel[i] > idxToDel[j]
		})
		for _, idx := range idxToDel {
			input = remove(input, idx)
		}
	}

	return wires
}

func andArguments(line string, wires map[string]uint16) (uint16, uint16, string, bool) {
	arguments := strings.Split(line, " AND ")
	first := arguments[0]
	second := strings.Split(arguments[1], " -> ")[0]
	result := strings.Split(arguments[1], " -> ")[1]
	_, fExist := wires[first]
	_, sExist := wires[second]
	if fExist && sExist {
		return wires[first], wires[second], result, true
	} else {
		if fExist {
			secondNum, err := strconv.Atoi(second)
			if err == nil {
				return wires[first], uint16(secondNum), result, true
			}
		} else if sExist {
			firstNum, err := strconv.Atoi(first)
			if err == nil {
				return uint16(firstNum), wires[second], result, true
			}
		}
	}
	return 0, 0, "0", false
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
