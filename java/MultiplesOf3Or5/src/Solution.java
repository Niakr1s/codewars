import java.util.stream.Stream;

public class Solution {
    public int solution(int number) {
        return Stream.iterate(0, i -> ++i).limit(number)
                .filter(i -> i % 3 == 0 || i % 5 == 0).reduce(0, Integer::sum);
    }
}