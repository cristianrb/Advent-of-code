package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	println(lowestMd5Number("iwrupvqb", 5))
	println(lowestMd5Number("iwrupvqb", 6))
}

func lowestMd5Number(secretKey string, nZeros int) int {
	found := false
	res := 0
	for !found {
		data := []byte(secretKey + strconv.Itoa(res))
		hexRes := fmt.Sprintf("%x", md5.Sum(data))
		firstNChars, err := strconv.Atoi(hexRes[0:nZeros])
		if firstNChars == 0 && err == nil {
			found = true
		}
		res++
	}

	return res - 1
}
