package ContainerWithMostWater


//[1,8,6,2,5,4,8,3,7]
//暴力解法
//func maxArea(height []int) int {
//	l := len(height)
//	result := 0
//	for i := 2; i <= l; i++ {
//		for j := 1; j <= i; j++ {
//			result = max(result, (i - j)*min(height[i - 1], height[j - 1]))
//		}
//	}
//	return result
//}



//双指针
func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	result := 0
	for i < j {
		result = max(result, min(height[i], height[j])*(j - i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}


func min(v, u int) int {
	if v < u {
		return v
	}
	return u
}


func max(v, u int) int {
	if v > u {
		return v
	}
	return u
}