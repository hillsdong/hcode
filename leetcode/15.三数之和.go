package leetcode

import (
	"sort"
)

// 夹逼 优化
func threeSum(nums []int) [][]int {
	ret := [][]int{}
	l := len(nums)
	if l < 3 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, l-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k > j && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			if k > j && nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[k], nums[j]})
			}
		}
	}
	return ret

}

// 夹逼
func threeSum4(nums []int) [][]int {
	ret := [][]int{}
	l := len(nums)
	if l < 3 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				res := []int{nums[i], nums[k], nums[j]}
				ret = append(ret, res)

				j++
				if k != j {
					k--
				}
			}
		}
	}
	return ret

}

// 哈希
func threeSum3(nums []int) [][]int {
	ret := [][]int{}
	l := len(nums)
	if l < 3 {
		return ret
	}

	sort.Ints(nums)
	sameHash := map[int]map[int]map[int]bool{}

	for i := 0; i < l-2; i++ {
		numsHash := map[int]int{}
		for j := i + 1; j < l; j++ {
			if k, ok := numsHash[0-nums[i]-nums[j]]; ok {
				res := []int{nums[i], nums[k], nums[j]}
				if !same(sameHash, res) {
					ret = append(ret, res)
				}
			}
			numsHash[nums[j]] = j
		}
	}

	return ret
}

// 暴力
func threeSum2(nums []int) [][]int {
	ret := [][]int{}
	l := len(nums)
	if l < 3 {
		return ret
	}

	sort.Ints(nums)
	sameHash := map[int]map[int]map[int]bool{}

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res := []int{nums[i], nums[j], nums[k]}
					if !same(sameHash, res) {
						ret = append(ret, res)
					}
				}
			}
		}
	}

	return ret
}

func same(sameHash map[int]map[int]map[int]bool, res []int) bool {
	if v, ok := sameHash[res[0]]; ok {
		if v, ok := v[res[1]]; ok {
			if _, ok := v[res[2]]; ok {
				return true
			} else {
				sameHash[res[0]][res[1]][res[2]] = true
			}
		} else {
			sameHash[res[0]][res[1]] = map[int]bool{}
			sameHash[res[0]][res[1]][res[2]] = true
		}
	} else {
		sameHash[res[0]] = map[int]map[int]bool{}
		sameHash[res[0]][res[1]] = map[int]bool{}
		sameHash[res[0]][res[1]][res[2]] = true
	}

	return false
}
