package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	people := make([][3]int, 0)
	for i := 0; i < n; i++ {
		var t, z, y int
		fmt.Scan(&t, &z, &y)
		people = append(people, [3]int{t, z, y})
	}
	left := 0
	right := 15000 * 2
	for left < right {
		mid := (left + right) / 2
		if qsharikov(people, mid) >= m {
			right = mid
		} else {
			left = mid + 1
		}
	}
	anstime := left
	ans := qsharikov2(people, anstime, m)
	fmt.Println(left)
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
	}
}

func qsharikov(people [][3]int, tplan int) int {
	var ans int
	for _, value := range people {
		var count, tactual, cycle int
		t, y, z := value[0], value[1], value[2]
		flag := true
		if tactual+t > tplan {
			flag = false
		}
		for flag {
			for (count == cycle*y || count-(cycle*y) != 0) && flag {
				count += 1
				tactual += t
				if tactual+t > tplan {
					flag = false
				}
			}
			if tplan >= tactual+z+t {
				tactual += z
				cycle += 1
			} else {
				flag = false
			}
		}
		ans += count
	}
	return ans
}

func qsharikov2(people [][3]int, tplan int, shariki int) []int {
	effectivity := make([]float32, 0)
	counts := make([]int, 0)
	cycles := make([]int, 0)
	var sum int
	for _, value := range people {
		var count, tactual, cycle int
		t, y, z := value[0], value[1], value[2]
		flag := true
		for flag {
			for (count == cycle*y || count-(cycle*y) != 0) && flag {
				count += 1
				tactual += t
				if tactual+t > tplan {
					flag = false
				}
			}
			if tplan >= tactual+z+t {
				tactual += z
				cycle += 1
			} else {
				flag = false
			}
		}
		effectivity = append(effectivity, float32(t)+float32((z*cycle))/float32(count))
		counts = append(counts, count)
		cycles = append(cycles, cycle)
		sum += count
	}
	for sum > shariki {
		index := findmax(effectivity)
		t, y, z := people[index][0], people[index][1], people[index][2]
		minus := counts[index] - cycles[index]*y
		if minus == 1 {
			cycles[index]--
			counts[index]--
			sum--
			effectivity[index] = float32(t) + float32((z*cycles[index]))/float32(counts[index])
		} else {
			counts[index]--
			sum--
			effectivity[index] = float32(t) + float32((z*cycles[index]))/float32(counts[index])
		}
	}
	return counts
}

func findmax(effectivity []float32) int {
	max := effectivity[0]
	maxind := 0
	for index, value := range effectivity {
		if value > max {
			max = value
			maxind = index
		}
	}
	return maxind
}
