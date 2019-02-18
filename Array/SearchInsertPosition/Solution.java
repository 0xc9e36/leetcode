package SearchInsertPosition;

class Solution {
    public int searchInsert(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;

        if (target > nums[right]) {
            return right + 1;
        } else if (target < nums[left]) {
            return left;
        }

        int mid = 0;
        while(left <= right) {
            mid = left + (right - left)/2;
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] > target) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        if (target > nums[mid]) {
            return mid + 1;
        }
        return mid;
    }
}