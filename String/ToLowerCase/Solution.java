package ToLowerCase;

public class Solution {
    public String toLowerCase(String str) {
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < str.length(); ++i){
            char c = str.charAt(i);
            if (c >= 'A' && c <= 'Z') {
                c = (char)(c + 32);
            }
            sb.append(c);
        }
        return sb.toString();
    }
}

func toLowerCase(str string) string {
    res := ""
    for i := range str {
        c := str[i]
        if c >= 'A' && c<= 'Z' {
            c += 32
        }
        res += string(c)
    }
    return res
}
