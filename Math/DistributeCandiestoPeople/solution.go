package DistributeCandiestoPeople



func distributeCandies(candies int, num_people int) []int {

	loop, result := 1, make([]int, num_people)
	for candies >= 0 {
		for i := 0; i < num_people; i++ {
			need :=  i + 1 + num_people * (loop - 1)
			if candies < need {
				result[i] +=  candies
				return result
			}
			result[i] += need
			candies -= need

		}
		loop++
	}
	return result
}
