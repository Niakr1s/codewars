import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class DirReduction {
    private static final List<List<String>> pairs = List.of(List.of("SOUTH", "NORTH"), List.of("WEST", "EAST"));

    public static String[] dirReduc(String[] arr) {
        List<String> list = new LinkedList<>(Arrays.asList(arr));
        while (clean(list)) {
        }
        return list.toArray(String[]::new);
    }

    private static boolean clean(List<String> list) {
        for (int i = 0; i < list.size() - 1; i++) {
            List<String> pair = list.subList(i, i + 2);
            if (isOpposite(pair)) {
                list.remove(i);
                list.remove(i);
                return true;
            }
        }
        return false;
    }

    private static boolean isOpposite(List<String> pair) {
        for (List<String> oppositePair : pairs) {
            if (pair.containsAll(oppositePair)) return true;
        }
        return false;
    }
}
