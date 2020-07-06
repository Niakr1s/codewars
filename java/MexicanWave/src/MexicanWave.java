import java.util.stream.IntStream;

public class MexicanWave {
    public static String[] wave(String str) {
        return IntStream.range(0, str.length()).mapToObj(i -> capitalize(str, i))
                .filter(s -> !s.equals(str))
                .toArray(String[]::new);
    }

    private static String capitalize(String str, int pos) {
        return str.substring(0, pos) + str.substring(pos, pos + 1).toUpperCase() +
                str.substring(pos + 1);
    }
}