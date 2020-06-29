public class StringSplit {
    public static String[] solution(String s) {
        if (s.length() % 2 != 0) s = s + "_";
        String[] res = new String[s.length() / 2];
        for (int i = 0; i < res.length; i++) {
            res[i] = s.substring(i*2, i*2+2);
        }
        return res;
    }
}