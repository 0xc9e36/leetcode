package MaximumSubarray;

/**
 * Solution class
 *
 * @author tanwei
 * @date 17-11-30
 */
class Solution {
    public int maxSubArray(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        int[] sum = new int[nums.length];
        sum[0] = nums[0];
        int max = sum[0];
        for(int i = 1; i < sum.length; ++i) {
            sum[i] = Math.max(sum[i - 1] + nums[i], nums[i]);
            max = Math.max(max, sum[i]);
        }
        return max;
    }
}