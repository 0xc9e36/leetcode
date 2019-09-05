package RemoveElement;

class Solution {
    public int removeElement(int[] nums, int val) {
        int j = 0;
        for(int i = 0; i < nums.length; ++i) {
            if (nums[i] != val) {
                nums[j++] = nums[i];
            }
        }
        return j;
    }
}


//golang
func removeElement(nums []int, val int) int {
    result := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[result] = nums[i]
            result++
        }
    }
    return result
}