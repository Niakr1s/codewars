class Solution {
    static int stray(int[] numbers) {
        int repeatingNumber = getRepeatingNumber(numbers);
        for (int num : numbers) {
            if (num != repeatingNumber) {
                return num;
            }
        }
        return 0;
    }

    private static int getRepeatingNumber(int[] numbers) {
        if (numbers[0] == numbers[1] || numbers[0] == numbers[2]) return numbers[0];
        return numbers[2];
    }
}