import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.List;
import java.util.Stack;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Solution {

    public static void main(String[] args) throws IOException {

        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        String[] firstLine = bufferedReader.readLine().replaceAll("\\s+$", "").split(" ");

        Stack<Integer> s = new Stack<>();
        Stack<Integer> maxStack = new Stack<>();

        int limit = Integer.parseInt(firstLine[0]);

        for (int i = 0; i < limit; i++) {
            List<Integer> querry = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                    .map(Integer::parseInt)
                    .collect(Collectors.toList());

            if (querry.get(0) == 1){
                int item = querry.get(1);
                s.push(item);
                if (maxStack.isEmpty() || item >= maxStack.peek()) {
                    maxStack.push(item);
                }
            } else if (querry.get(0) == 2) {
                int item = s.pop();
                if (!maxStack.isEmpty() && item == maxStack.peek()){
                    maxStack.pop();
                }
            } else {
                if (!maxStack.isEmpty()){
                    System.out.println(maxStack.peek());    
                }
                
            }
        }

    }

}