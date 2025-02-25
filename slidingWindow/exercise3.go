package slidingwindow

//### **3. Longest Substring with K Distinct Characters**
//**Problem:**
//Given a string, find the length of the longest substring that contains at most `K` distinct characters.
//
//**Example:**
//```plaintext
//Input: str = "araaci", K = 2
//Output: 4
//Explanation: The longest substring with ≤ 2 distinct characters is "araa".
//```
//**Sliding Window Approach:**
//1. Expand the window (`end` pointer) while tracking the number of distinct characters.
//2. If distinct characters exceed `K`, shrink the window (`start` pointer).
//3. Keep track of the maximum substring length.
//
//---
