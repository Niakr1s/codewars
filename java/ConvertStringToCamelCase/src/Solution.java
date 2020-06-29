class Solution{
    static String toCamelCase(String s){
        String[] subStrings = s.split("[^a-zA-Z]");
        // converting all substrings to uppercase, skipping first substring
        for (int i = 1; i < subStrings.length; i++) {
            String str = subStrings[i];
            subStrings[i] = Character.toUpperCase(str.charAt(0)) + str.substring(1);
        }
        return String.join("", subStrings);
    }
}