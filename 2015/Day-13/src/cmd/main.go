package main

import (
	"cristianrb2015/io"
	"cristianrb2015/util"
	"strconv"
	"strings"
)

type Relationship struct {
	Person1        string
	Person2        string
	HappinessUnits int
}

func main() {
	lines, _ := io.ReadLines("Day-13/inputs/input_01")
	// Sol 1
	persons, relations := generatePersonsAndRelationships(lines, false)
	permutations := generatePermutations(persons)
	allHappiness := findTotalHappiness(permutations, relations)
	println(util.MaxOf(allHappiness...))

	// Sol 2
	persons, relations = generatePersonsAndRelationships(lines, true)
	permutations = generatePermutations(persons)
	allHappiness = findTotalHappiness(permutations, relations)
	println(util.MaxOf(allHappiness...))
}

func findTotalHappiness(permutations [][]string, relations []Relationship) []int {
	var allHappiness []int
	for _, persons := range permutations {
		lastPerson := persons[0]
		totalHappiness := getRelationship(relations, lastPerson, persons[len(persons)-1])
		for idx := 1; idx < len(persons); idx++ {
			actualPerson := persons[idx]
			totalHappiness += getRelationship(relations, lastPerson, actualPerson)
			totalHappiness += getRelationship(relations, actualPerson, lastPerson)
			lastPerson = actualPerson
		}
		totalHappiness += getRelationship(relations, lastPerson, persons[0])
		allHappiness = append(allHappiness, totalHappiness)
	}
	return allHappiness
}

func getRelationship(relations []Relationship, person1 string, person2 string) int {
	for _, relation := range relations {
		if relation.Person1 == person1 && relation.Person2 == person2 {
			return relation.HappinessUnits
		}
	}
	return 0
}

func generatePermutations(persons []string) [][]string {
	var permutations [][]string
	util.Perm(persons, func(persons []string) {
		var permutation []string
		for _, person := range persons {
			permutation = append(permutation, person)
		}
		permutations = append(permutations, permutation)
	})
	return permutations
}

func generatePersonsAndRelationships(lines []string, addMe bool) ([]string, []Relationship) {
	var relations []Relationship
	var persons []string
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		person1 := lineSplit[0]
		person2 := lineSplit[len(lineSplit)-1]
		person2 = person2[0 : len(person2)-1]
		operation := lineSplit[2]
		hapinnessUnits, _ := strconv.Atoi(lineSplit[3])

		if operation == "lose" {
			hapinnessUnits *= -1
		}

		if !util.Contains(persons, person1) {
			persons = append(persons, person1)
		}
		relations = append(relations, Relationship{
			person1,
			person2,
			hapinnessUnits,
		})

	}

	if addMe {
		for _, person := range persons {
			relations = append(relations, Relationship{
				person,
				"Cristian",
				0,
			})
			relations = append(relations, Relationship{
				"Cristian",
				person,
				0,
			})
		}
		persons = append(persons, "Cristian")
	}

	return persons, relations
}
