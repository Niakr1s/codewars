public class Kata {

    public static String encrypt(String text, final int n) {
        if (text == null) return null;
        for (int i = 0; i < n; i++) {
            text = encryptOnce(text);
        }
        return text;
    }

    private static String encryptOnce(final String text) {
        StringBuilder builder = new StringBuilder();
        for (int i = 1; i < text.length(); i += 2) {
            builder.append(text.charAt(i));
        }
        for (int i = 0; i < text.length(); i += 2) {
            builder.append(text.charAt(i));
        }
        return builder.toString();
    }

    public static String decrypt(String text, final int n) {
        if (text == null) return null;
        for (int i = 0; i < n; i++) {
            text = decryptOnce(text);
        }
        return text;
    }

    private static String decryptOnce(String text) {
        StringBuilder builder = new StringBuilder();
        final int firstPartLength = text.length() / 2;
        for (int left = 0, right = firstPartLength; right < text.length(); left++, right++) {
            builder.append(text.charAt(right));
            if (left < firstPartLength) builder.append(text.charAt(left));
        }
        return builder.toString();
    }
}