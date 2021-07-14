package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	//	list := []int{-500, -481, -480, -469, -437, -423, -408, -403, -397, -381, -379, -377, -353, -347, -337, -327, -313, -307, -299, -278, -265, -258, -235, -227, -225, -193, -192, -177, -176, -173, -170, -164, -162, -157, -147, -118, -115, -83, -64, -46, -36, -35, -11, 0, 0, 33, 40, 51, 54, 74, 93, 101, 104, 105, 112, 112, 116, 129, 133, 146, 152, 157, 158, 166, 177, 183, 186, 220, 263, 273, 320, 328, 332, 356, 357, 363, 372, 397, 399, 420, 422, 429, 433, 451, 464, 484, 485, 498, 499}
	list := []int{1, 0, -1, 0, -2, 2}
	target := 0

	ans := fourSum(list, target)

	fmt.Println(ans)
}

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	cnt := len(nums)

	if cnt < 4 {
		return ans
	}

	for i := 0; i < cnt; i++ {
		a := nums[i]
		for j := i + 1; j < cnt; j++ {
			b := nums[j]
			for k := j + 1; k < cnt; k++ {
				c := nums[k]
				for l := k + 1; l < cnt; l++ {
					d := nums[l]
					if (a + b + c + d) == target {
						n := []int{a, b, c, d}
						if !contains(ans, n) {
							ans = append(ans, n)
						}
					}
				}
			}
		}
	}

	return ans
}

func contains(list [][]int, nums []int) bool {
	sort.Ints(nums)
	for _, v := range list {
		sort.Ints(v)
		if reflect.DeepEqual(v, nums) {
			return true
		}
	}
	return false
}
