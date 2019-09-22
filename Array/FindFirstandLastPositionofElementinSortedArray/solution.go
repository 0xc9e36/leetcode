package FindFirstandLastPositionofElementinSortedArray

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums) - 1
	res := make([]int, 2)
	res[0], res[1] = -1, -1

	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if target <= nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if lo >= len(nums) || nums[lo] != target {
		return res
	}

	res[0] = lo

	lo, hi = 0, len(nums) - 1
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if target >=  nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	res[1] = hi
	return res
}
