package ReverseWordsinaStringIII;


public class Solution {
    public String reverseWords(String s) {
        String[] array = s.split(" ");
        int right = array.length;
        StringBuilder res = new StringBuilder();
        for(int i = 0; i < right; ++i) {
            res.append(new StringBuilder(array[i]).reverse());
            if (i < right - 1) {
                res.append(" ");
            }
        }
        return res.toString();
    }
}