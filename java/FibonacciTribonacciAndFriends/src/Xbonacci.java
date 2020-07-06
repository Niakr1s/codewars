import java.util.Arrays;

public class Xbonacci {

    public double[] xbonacci(double[] signature, int n) {
        double[] res = Arrays.copyOf(signature, n);
        for (int i = signature.length; i < n; i++) {
            res[i] = sum(res, i - signature.length, i);
        }
        return res;
    }

    private double sum(double[] arr, int from, int to) {
        double res = 0;
        for (int i = from; i < to; i++) res += arr[i];
        return res;
    }
}