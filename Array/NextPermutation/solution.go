package NextPermutation

import "sort"

func nextPermutation(nums []int)  {

	j, l := len(nums) - 1, len(nums) - 1

	if l <= 0 {
		return
	}

	//最后两位从大到小排列  ... 1 5, 直接交换
	if nums[j] > nums[j - 1] {
		nums[j], nums[j - 1] = nums[j - 1], nums[j]
		return
	}

	for j >= 1 && nums[j] <= nums[j - 1] {
		j--
	}

	//从大到小排列  5 4 3 2 1 ==>  1 2 3 4 5
	if j == 0 {
		sort.Ints(nums)
		return
	}


	//2 5 3 1   找出右边最接近 j - 1号元素交换
	k := l
	for ; k >= j + 1; k-- {
		if nums[k] > nums[j - 1] {
			break
		}
	}
	nums[j - 1], nums[k] = nums[k], nums[j - 1]

	sort.Ints(nums[j:])

	return
}