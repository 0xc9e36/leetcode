package RobotReturntoOrigin;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public boolean judgeCircle(String moves) {
        Map<Character, Integer> map = new HashMap<>();

        for(int i = 0; i < moves.length(); ++i) {
            char c = moves.charAt(i);
            map.merge(c, 1, Integer::sum);
        }
        int l = map.containsKey('L') ? map.get('L') : 0;
        int r = map.containsKey('L') ? map.get('R') : 0;
        int u = map.containsKey('L') ? map.get('U') : 0;
        int d = map.containsKey('L') ? map.get('D') : 0;
        return l == r && u == d;
    }
}