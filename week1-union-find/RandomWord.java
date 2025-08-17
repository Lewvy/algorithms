import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class RandomWord {
	public static void main(String[] args) {
		String str = "", winner = "";
		int i = 0;
		while (!StdIn.isEmpty()) {
			str = StdIn.readString();
			i++;
			if (StdRandom.bernoulli((1.0 / i))) {
				winner = str;
			}
		}
        StdOut.println(winner);
	}
}
