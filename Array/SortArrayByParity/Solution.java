package SortArrayByParity;

class Solution {
    public int[] sortArrayByParity(int[] A) {
        if (A.length <= 0) {
            return A;
        }
        int left = 0;
        int right = A.length - 1;
        while (left < right) {
            while (A[left] % 2 == 0) {
                left++;
                if (left >= A.length) {
                    break;
                }
            }
            while (A[right] % 2 == 1) {
                right--;
                if (right <= 0) {
                    break;
                }
            }
            if (left < right) {
                int temp = A[left];
                A[left] = A[right];
                A[right] = temp;
                left++;
                right--;
            }
        }
        return A;
    }
}