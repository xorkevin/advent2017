package main

import (
	"fmt"
	"strconv"
)

const (
	puzzleInput = "xlqgujun"
)

var (
	part2Suffix = []int{17, 31, 73, 47, 23}
)

func main() {
	fmt.Println(diskUsage(puzzleInput))
	fmt.Println(regions(genGrid(puzzleInput)))
}

func regions(grid [][]int) int {
	k := 0
	for r, i := range grid {
		for c, j := range i {
			if j < 0 {
				k++
				findRegion(grid, r, c, k)
			}
		}
	}
	return k
}

func findRegion(grid [][]int, r, c, k int) {
	l := len(grid)
	l2 := len(grid[0])
	openList := [][]int{[]int{r, c}}
	for len(openList) > 0 {
		point := openList[0]
		grid[point[0]][point[1]] = k
		openList = openList[1:]
		n := getNeighbors(point[0], point[1], l, l2)
		for _, i := range n {
			if grid[i[0]][i[1]] < 0 {
				openList = append(openList, i)
			}
		}
	}
}

func getNeighbors(r, c, sy, sx int) [][]int {
	neighbors := [][]int{}
	if r > 0 {
		neighbors = append(neighbors, []int{r - 1, c})
	}
	if r < sy-1 {
		neighbors = append(neighbors, []int{r + 1, c})
	}
	if c > 0 {
		neighbors = append(neighbors, []int{r, c - 1})
	}
	if c < sx-1 {
		neighbors = append(neighbors, []int{r, c + 1})
	}
	return neighbors
}

func genGrid(hashInput string) [][]int {
	grid := make([][]int, 0, 128)
	for i := 0; i < 128; i++ {
		gridRow := make([]int, 128)
		nums := hash(hashInput + "-" + strconv.Itoa(i))
		for n, num := range nums {
			for j := 0; j < 8; j++ {
				gridRow[8*n+7-j] = -num % 2
				num /= 2
			}
		}
		grid = append(grid, gridRow)
	}
	return grid
}

func diskUsage(hashInput string) int {
	k := 0
	for i := 0; i < 128; i++ {
		k += hashUsage(hash(hashInput + "-" + strconv.Itoa(i)))
	}
	return k
}

func hashUsage(nums []int) int {
	k := 0
	for _, i := range nums {
		for i > 0 {
			k += i % 2
			i /= 2
		}
	}
	return k
}

func hash(hashInput string) []int {
	nums := sequence(hashInput)

	list := make([]int, 0, 256)
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	k := 0
	skip := 0

	for i := 0; i < 64; i++ {
		for _, i := range nums {
			flip(list, k, i)
			k = (skip + k + i) % len(list)
			skip++
		}
	}

	return dense(list)
}

func flip(nums []int, start, size int) {
	l := len(nums)
	end := (l + start + size - 1) % l
	for i := 0; i < size/2; i++ {
		swap(nums, (start+i)%l, (l+end-i)%l)
	}
}

func swap(nums []int, a, b int) {
	k := nums[a]
	nums[a] = nums[b]
	nums[b] = k
}

func sequence(inp string) []int {
	binp := []byte(inp)
	nums := make([]int, 0, len(binp)+len(part2Suffix))
	for _, i := range binp {
		nums = append(nums, int(i))
	}
	nums = append(nums, part2Suffix...)
	return nums
}

func dense(nums []int) []int {
	list := make([]int, 0, len(nums)/16)
	for i := 0; i < len(nums); i += 16 {
		k := 0
		for j := 0; j < 16; j++ {
			k ^= nums[i+j]
		}
		list = append(list, k)
	}
	return list
}
