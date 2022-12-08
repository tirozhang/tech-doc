package Rabbit_Count

import "fmt"

// n是天数 k繁殖周期 d 为死亡天数
func rabbitCount(n, k, d int) int {
	rabbitList := []int{d}

	fmt.Println(rabbitList)
	count := 0

	return count * 2
}

// n当前天数 k繁殖周期
func increase(rabbits []int, n, k, d int) []int {
	for _, v := range rabbits {
		if v > n && (n-v+1-k+1)%k == 0 {
			rabbits = append(rabbits, n+d-1)
		}
	}
	return rabbits
}
