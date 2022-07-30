package main

import (
	"bytes"
	"fmt"
)

func main() {
	result := "1113122113"
	for i := 0; i < 60; i++ {
		result = lookAndSay(result)
	}
	println(len(result))
}

func lookAndSay(s string) string {
	actualNum := s[0]
	actualCount := 1
	var buffer bytes.Buffer
	for i := 1; i < len(s); i++ {
		if actualNum == s[i] {
			actualCount++
		} else {
			buffer.WriteString(fmt.Sprintf("%d%s", actualCount, string(actualNum)))
			actualCount = 1
			actualNum = s[i]
		}
	}
	buffer.WriteString(fmt.Sprintf("%d%s", actualCount, string(actualNum)))

	return buffer.String()
}
