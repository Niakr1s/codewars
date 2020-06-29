import javax.print.DocFlavor;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class Diamond {
    public static String print(int n) {
        if (n < 0 || n % 2 == 0) return null;
        int middle = n / 2;
        List<String> list = new ArrayList<>();
        for (int i = 0; i < middle; i++) {
            list.add(getLine(i * 2 + 1, n));
        }
        list.add(getLine(n, n));
        addBottomOfDiamond(list);
        return String.join("", list);
    }

    /** Last line of list should be middle line of diamond (eg size of list should be odd). */
    private static void addBottomOfDiamond(List<String> list) {
        List<String> part = new ArrayList<>(list.subList(0, list.size() - 1));
        Collections.reverse(part);
        list.addAll(part);
    }

    private static String getLine(int amountOfAsterisks, int maxStringLength) {
        int whitespaces = (maxStringLength - amountOfAsterisks) / 2;
        return " ".repeat(whitespaces) + "*".repeat(amountOfAsterisks) + "\n";
    }
}