import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class DescendingOrder {
    public static int sortDesc(final int num) {
        List<Integer> digits = toDigits(num);
        digits.sort(IntComparators.ascending());
        return toNumber(digits);
    }

    private static List<Integer> toDigits(int num) {
        List<Integer> list = new ArrayList<>();
        while (num != 0) {
            int digit = num % 10;
            list.add(digit);
            num = num / 10;
        }
        return list;
    }

    private static int toNumber(List<Integer> list) {
        int multiplier = 1;
        int res = 0;
        for (int i = list.size() - 1; i >= 0; i--) {
            res += list.get(i) * multiplier;
            multiplier *= 10;
        }
        return res;
    }
}

class IntComparators {
    static Comparator<Integer> ascending() {
        return (o1, o2) -> o2 - o1;
    }
}