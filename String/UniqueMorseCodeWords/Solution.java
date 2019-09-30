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



func uniqueMorseRepresentations(words []string) int {
    code := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

    m := map[string]struct{}{}

    res := 0
    for _, word := range words {
        t := ""
        for j := range word {
            t += code[word[j] - 97]
        }
        _, ok := m[t]
        if !ok {
            res++
            m[t] = struct{}{}
        }
    }
    return res
}