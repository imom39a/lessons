class Solution {
    
    public static void main(String[] args) {

        int result = searchInsert(new int[]{1, 2, 3, 4, 5, 6}, 4);
        System.out.println(result);

    }

    public static int searchInsert(int[] nums, int target) {

        int lPtr = 0;
        int rPtr = nums.length - 1;

        while (lPtr <= rPtr) {
            int mPtr = (rPtr + lPtr) / 2;
            if (nums[mPtr] == target) {
                return mPtr;
            } else if (target < nums[mPtr]) {
                rPtr = mPtr - 1;
            } else {
                lPtr = mPtr + 1;
            }
        }

        return lPtr;
    }

}