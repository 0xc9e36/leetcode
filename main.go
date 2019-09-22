package main

func main() {
	r := []int{5,7,7,8,8,10}
	searchRange(r, 8)
}

//5,7,7,8,8,10
//8

func searchRange(nums []int, target int) []int {

	lo, hi := 0, len(nums) - 1
	res := make([]int ,2)
	res[0], res[1] = -1, -1
	if hi < lo {
		return res
	}

	// left
	//
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if hi == len(nums) - 1 || nums[hi] != target {
		return res
	}

	res[0] = hi
	hi = len(nums) - 1

	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if target >= nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	res[1] = hi - 1
	return res
 }