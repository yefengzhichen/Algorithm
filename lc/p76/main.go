package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDic(t string) map[byte]int {
	resDic := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		if _, ok := resDic[t[i]]; ok {
			resDic[t[i]] += 1
		} else {
			resDic[t[i]] = 1
		}
	}
	return resDic
}

func minWindow(s string, t string) string {
	stack := make([]int, 0)
	dic, origDic, res := getDic(t), getDic(t), ""
	curDic := make(map[byte]int, 0)
	s += " "
	for i := 0; i < len(s); i++ {
		stack = append(stack, i)
		curDic[s[i]] += 1
		if len(dic) > 0 {
			v, ok := dic[s[i]]
			if v == 1 {
				delete(dic, s[i])
			} else if ok {
				dic[s[i]] -= 1
			}
		}
		if len(dic) == 0 {
			for len(stack) > 0 && curDic[s[stack[0]]] > origDic[s[stack[0]]] {
				start := stack[0]
				stack = stack[1:]
				curDic[s[start]] -= 1
			}
			tmp := s[stack[0] : stack[len(stack)-1]+1]
			if len(res) == 0 || len(tmp) < len(res) {
				res = tmp
			}
		}
	}
	return res
}

func main() {
	fmt.Println("a")
}
