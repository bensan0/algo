package main

import "fmt"

func main() {
	fmt.Println(fibonacci(9))

	fmt.Println(twosum([]int{3, 3}, 6))
}

func fibonacci(n int) int {
	a := 0
	b := 1

	if n == 0 {
		return a
	} else if n == 1 {
		return b
	}

	for i := 1; i < n; i++ {
		temp := a
		a = b
		b = temp + b
	}
	return b
}

func twosum(arr []int, target int) []int {

	m := make(map[int]int)
	for i, v := range arr {
		diff := target - v
		if ani, ok := m[diff]; ok {
			if ani == i {
				continue
			} else {
				return []int{i, ani}
			}
		}
		m[v] = i
	}
	return []int{0, 0}

}
