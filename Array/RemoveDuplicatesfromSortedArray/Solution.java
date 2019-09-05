package RemoveDuplicatesfromSortedArray;

class Solution {
    public int removeDuplicates(int[] nums) {
        int left = 0;
        int right = 1;
        for(int i = 0; i < nums.length && right < nums.length; ++i) {
            if (nums[left] != nums[right]) {
                nums[left + 1] = nums[right];
                left++;
            }
            right++;
        }
        return left + 1;
    }
}



// golang
func removeDuplicates(nums []int) int {

    i := 1
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[j - 1] {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}

