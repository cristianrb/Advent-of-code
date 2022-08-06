package main

import (
	"cristianrb2015/io"
	"cristianrb2015/util"
	"fmt"
	"sort"
)

type Reeinder struct {
	name              string
	topSpeed          int
	flySeconds        int
	restSeconds       int
	restSecondsLeft   int
	timeUntilNextRest int
	distance          int
	score             int
}

func main() {
	lines, _ := io.ReadLines("Day-14/inputs/input_01")
	reeinders := findReeinderStats(lines)
	distances, scores := compute(reeinders, 2503)
	// Sol 1
	println(util.MaxOf(distances...))
	// Sol 2
	println(util.MaxOf(scores...))
}

func compute(reeinders []*Reeinder, seconds int) ([]int, []int) {
	for seconds > 0 {
		for _, reeinder := range reeinders {
			if reeinder.restSecondsLeft > 0 {
				reeinder.restSecondsLeft -= 1

			} else {
				reeinder.distance += reeinder.topSpeed
				reeinder.timeUntilNextRest -= 1
				if reeinder.timeUntilNextRest == 0 {
					reeinder.restSecondsLeft = reeinder.restSeconds
					reeinder.timeUntilNextRest = reeinder.flySeconds
				}
			}
		}

		addScoreIfLeading(reeinders)
		seconds--
	}

	var distances, scores []int
	for _, reeinder := range reeinders {
		distances = append(distances, reeinder.distance)
	}
	for _, reeinder := range reeinders {
		scores = append(scores, reeinder.score)
	}

	return distances, scores
}

func addScoreIfLeading(reeinders []*Reeinder) {
	sort.Slice(reeinders, func(i, j int) bool {
		return reeinders[i].distance > reeinders[j].distance
	})
	maxDistance := reeinders[0].distance
	for _, reeinder := range reeinders {
		if reeinder.distance == maxDistance {
			reeinder.score++
		}
	}
}

func findReeinderStats(lines []string) []*Reeinder {
	var reeinders []*Reeinder
	for _, line := range lines {
		var name string
		var topSpeed, flySeconds, restSeconds int
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &topSpeed, &flySeconds, &restSeconds)

		reeinder := &Reeinder{
			name,
			topSpeed,
			flySeconds,
			restSeconds,
			0,
			flySeconds,
			0,
			0,
		}
		reeinders = append(reeinders, reeinder)
	}
	return reeinders
}
