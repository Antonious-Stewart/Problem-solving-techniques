package slidingwindow

//### **2. Smallest Subarray with Sum Greater Than or Equal to `S`**
//**Problem:**
//Given an array of positive numbers and an integer `S`, find the length of the smallest contiguous subarray whose sum is greater than or equal to `S`. Return `0` if no such subarray exists.
//
//**Example:**
//Input: arr = [2, 1, 5, 2, 3, 2], S = 7
//Output: 2
//Explanation: The smallest subarray with sum ≥ 7 is [5, 2] (length 2).

//**Sliding Window Approach:**
//1. Use two pointers (`start` and `end`). Expand `end` until the sum is ≥ `S`.
//2. Shrink `start` to minimize the subarray size while keeping the sum ≥ `S`.
