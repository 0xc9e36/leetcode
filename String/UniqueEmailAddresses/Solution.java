package UniqueEmailAddresses;

import java.util.HashSet;
import java.util.Set;

class Solution {

    public int numUniqueEmails(String[] emails) {
        Set<String> set = new HashSet<>();
        for(String email : emails) {
            int index = email.indexOf('@');
            int plus = email.indexOf('+');
            String ele = null;
            if (plus != -1) {
                ele = email.substring(0, plus);
            } else {
                ele = email.substring(0, index);
            }
            set.add(ele.replace(".", "") + email.substring(index));
        }
        return set.size();
    }
}