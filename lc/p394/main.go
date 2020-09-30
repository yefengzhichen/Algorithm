package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	i := 0
	var res bytes.Buffer
	length := len(s)
	for i < length {
		if s[i] > '0' && s[i] < '9' {
			start := i
			for ; i < length && s[i] >= '0' && s[i] <= '9'; i++ {
			}
			num, _ := strconv.Atoi(s[start:i])
			start = i + 1
			stack := 0
			for i < length {
				if s[i] == '[' {
					stack++
				} else if s[i] == ']' {
					stack--
				}
				if stack == 0 {
					nested := decodeString(s[start:i])
					for l := 0; l < num; l++ {
						res.WriteString(nested)
					}
					i++
					break
				}
				i++
			}
		} else {
			start := i
			for ; i < length && (s[i] < '0' || s[i] > '9'); i++ {
			}
			res.WriteString(s[start:i])
		}
	}
	return res.String()
}

func main() {
	fmt.Println("a")
}
