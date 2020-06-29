import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class isogram {
    public static boolean isIsogram(String str) {
        if (str.isEmpty()) return true;
        str = str.toLowerCase();
        Set<String> set = new HashSet<>(List.of(str.split("")));
        return set.size() == str.length();
    }
}