package slidingwindow

//### **4. Longest Subarray with Sum ≤ `S`**
//**Problem:**
//Given an array of positive integers and a target sum `S`, find the length of the longest contiguous subarray with a sum ≤ `S`.
//
//**Example:**
//```plaintext
//Input: arr = [3, 1, 2, 5, 1, 1, 2, 3], S = 6
//Output: 4
//Explanation: The longest subarray is [3, 1, 2] or [1, 2, 2, 1], both with sum ≤ 6.
//```
//**Sliding Window Approach:**
//1. Expand `end` to increase the sum.
//2. If the sum exceeds `S`, shrink `start` to bring it back within limits.
//3. Track the maximum subarray length.
//
//---
