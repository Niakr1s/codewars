class Solution {
    static String toCamelCase(String s) {
        String[] subStrings = s.split("[^a-zA-Z]");
        if (subStrings.length == 0) return "";
        StringBuilder builder = new StringBuilder();
        builder.append(subStrings[0]);
        // converting all substrings to uppercase, skipping first substring
        for (int i = 1; i < subStrings.length; i++) {
            String str = subStrings[i];
            builder.append(Character.toUpperCase(str.charAt(0))).append(str.substring(1));
        }
        return builder.toString();
    }
}