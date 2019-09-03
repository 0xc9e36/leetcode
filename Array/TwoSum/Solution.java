package TwoSum;

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        if (nums.length < 2) {
            return new int[0];
        }

        Map<Integer, Integer> map = new HashMap();
        int[] res = new int[2];
        for (int i = 0; i < nums.length; ++i) {
            if (map.containsKey(target - nums[i])) {
                res[0] = map.get(target - nums[i]);
                res[1] = i;
                break;
            }
            map.put(nums[i], i);
        }
        return res;
    }
}









//golang


func twoSum(nums []int, target int) []int {
    result := make([]int, 2)
    hash := make(map[int]int, 0)
    for i, num := range nums {
        index, ok := hash[target - num]
        if ok {
            result[0] = index
            result[1] = i
            return result
        }
        hash[num] = i
    }
    return result
}