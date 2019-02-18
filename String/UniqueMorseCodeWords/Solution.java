package UniqueMorseCodeWords;

import java.util.HashSet;
import java.util.Set;

class Solution {
    public int uniqueMorseRepresentations(String[] words) {
        String[] array = new String[]{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
        Set<String> set = new HashSet<>();
        for(String word : words) {
            StringBuilder sb = new StringBuilder();
            String low = word.toLowerCase();
            for(int i = 0; i < low.length(); ++i) {
                sb.append(array[low.charAt(i) - 97]);
            }
            set.add(sb.toString());
        }
        return set.size();
    }
}