package slidingwindow

//### **5. Longest Substring Without Repeating Characters**
//**Problem:**
//Given a string, find the length of the longest substring without repeating characters.
//
//**Example:**
//```plaintext
//Input: str = "abcabcbb"
//Output: 3
//Explanation: The longest substring is "abc".
//```
//**Sliding Window Approach:**
//1. Use a set (or hashmap) to track characters in the current window.
//2. Expand `end` and add new characters to the set.
//3. If a duplicate is found, shrink `start` until all characters are unique.
//4. Keep track of the maximum substring length.
//
//---
