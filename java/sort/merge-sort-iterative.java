import java.util.Arrays;
import java.util.stream.Stream;

class Solution {

    public static void main(String[] args) {
        
       int[] result = mergeSort(new int[]{5,4,3,2,1,0});

      for (int i : result) {
           System.out.println(i);
       }
    }

    private static int[] mergeSort(int[] nums) {

        if (nums.length <= 1) {
            return nums;
        }

        int[] tmp = Arrays.copyOf(nums, nums.length);

        for (int subArrayBlockSize = 1; subArrayBlockSize < nums.length; subArrayBlockSize = subArrayBlockSize * 2) {
            //System.out.println("subArrayBlockSize " + subArrayBlockSize);
            for (int i = 0; i < nums.length - 1 ; i += subArrayBlockSize * 2) {
                
                int from = i;
                int mid = from + subArrayBlockSize - 1;
                int to = Math.min(((i + subArrayBlockSize * 2)),nums.length) - 1;
                                
                merge(nums, tmp, from, mid, to);
            }
        }

        return nums;
    }

    private static void merge(int[] nums, int[] tmp, int lIdx, int mIdx, int rIdx) {
        //System.out.println("calling merge " + "lIdx " + lIdx + " mIdx " + mIdx + " rIdx " + rIdx);
        int lPrt = lIdx;
        int rPrt = mIdx + 1;
        int tPrt = lIdx;

        while (lPrt <= mIdx && rPrt <= rIdx) {
            if (nums[lPrt] < nums[rPrt]) {
                tmp[tPrt++] = nums[lPrt++];
            } else {
                tmp[tPrt++] = nums[rPrt++];
            }
        }

        while (lPrt <= mIdx) {
            tmp[tPrt++] = nums[lPrt++];
        }

        while (rPrt <= rIdx) {
            tmp[tPrt++] = nums[rPrt++];
        }

        for (int i = lIdx; i <= rIdx; i++) {
            nums[i] = tmp[i];
        }

    }

    

}