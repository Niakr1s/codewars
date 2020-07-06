import java.util.Arrays;

public class Kata {
    public static int[] foldArray(int[] array, int runs) {
        if (runs == 0) return array;

        int[] res = Arrays.copyOf(array, (array.length + 1) / 2);
        int end = array.length % 2 == 0 ? res.length : res.length - 1;
        for (int i = 0; i < end; i++) {
            res[i] += array[array.length - 1 - i];
        }
        return foldArray(res, runs - 1);
    }
}