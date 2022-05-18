# LeetCode Programs
Adding to this readme is not very efficient, but gives me a moment to pause and collect myself.

## Arrays

| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[26](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)| <span style="color: green"> done </span>|[Remove Duplicates from Sorted Array](remove-duplicates-from-sorted-array.go) | Use read/write variables as the 2 pointers | 
|[118]( https://leetcode.com/problems/pascals-triangle/ )| <span style="color: green"> done </span>| [ Pascals Triangle ]( bits/pascals-triangle.go )| | 
|[122]( https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii )| <span style="color: orange "> redo </span>| [ Best Time to Buy and Sell Stock II ]( arrays/best-time-to-buy-and-sell-stock-ii.go )| | 
|[189]( https://leetcode.com/problems/rotate-array/ )| <span style="color: green "> done </span>| [ Rotate Array ]( arrays/rotate-array.go )| cyclic replacement is fun |
|[217]( https://leetcode.com/problems/contains-duplicate/ )| <span style="color: green "> done </span>| [ Contains Duplicate ]( arrays/contains-duplicate.go )| |
|[136]( https://leetcode.com/problems/single-number/ )| <span style="color: green "> done </span>| [ Single Number ]( arrays/single-number.go )| |
|[66]( https://leetcode.com/problems/plus-one/ )| <span style="color: green "> done </span>| [ Plus One ]( arrays/plus-one.go )| |
|[1]( https://leetcode.com/problems/two-sum/ )| <span style="color: green "> done </span>| [ Two Sum ]( arrays/two-sum.go )| |
|[36]( https://leetcode.com/problems/valid-sudoku/ )| <span style="color: green "> done </span>| [ Valid Sudoku ]( arrays/valid-sudoku.go )| |
|[48]( https://leetcode.com/problems/rotate-image/ )| <span style="color: red"> redo </span>| [ Rotate Image ]( arrays/rotate-image.go )| Redo from scratch |

## Strings 

| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[344]( https://leetcode.com/problems/reverse-string/ )| <span style="color: green"> done </span>| [ Reverse String ]( strings/reverse-string.go )| |
|[7]( https://leetcode.com/problems/reverse-integer/ )| <span style="color: green"> done </span>| [ Reverse Integer ]( strings/reverse-integer.go )| |
|[387]( https://leetcode.com/problems/first-unique-character-in-a-string/ )| <span style="color: green"> done </span>| [ First Unique Character in a String ]( strings/first-unique-character-in-a-string.go )| |
|[242]( https://leetcode.com/problems/valid-anagram/ )| <span style="color: green"> done </span>| [ Valid Anagram ]( strings/valid-anagram.go )| |
|[125]( https://leetcode.com/problems/valid-palindrome/ )| <span style="color: green"> done </span>| [ Valid Palindrome ]( strings/valid-palindrome.go )| |
|[28]( https://leetcode.com/problems/implement-strstr/ )| <span style="color: green"> done </span>| [ Implement strStr() ]( strings/implement-strStr.go )| |



### Lists, Stacks, Qs 

| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[20]( https://leetcode.com/problems/valid-parentheses/ )| <span style="color: green"> done </span>| [ Valid Parentheses ]( list_stk_q/valid-parentheses.go )| | 

## Trees

| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[104](https://leetcode.com/problems/maximum-depth-of-binary-tree/)| <span style="color: green"> done </span>| [Maximum Depth of Binary Tree ]( trees/maximum-depth-of-binary-tree.go ) | Try non-recursive to see how much faster it is | 
|[98](https://leetcode.com/problems/validate-binary-search-tree/)| <span style="color: green"> done </span>| [Validate Binary Search Tree]( trees/validate-binary-search-tree.go ) | There is a bottom up and top down approach to this, try the bottom up |
|[101](https://leetcode.com/problems/symmetric-tree/)| <span style="color: green"> done </span>| [ Symmetric Tree ]( trees/symmetric-tree.go )| |
|[102](https://leetcode.com/problems/binary-tree-level-order-traversal/)| <span style="color: green"> done </span>| [Binary Tree Level Order Traversal ]( trees/binary-tree-level-order-traversal.go )| |
|[108]( https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/ )| <span style="color: green"> done </span>| [ Convert Sorted Array to Binary Search Tree ]( trees/convert-sorted-array-to-binary-search-tree.go )| How would you do it iteratively? |

## Sorting Searching 
| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[268](https://leetcode.com/problems/missing-number/)| <span style="color: green"> done </span>| [ Missing Number ]( sorting/missing-number.go )| cycling sort is fun |

## Dynamic Programming 
| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[70]( https://leetcode.com/problems/climbing-stairs/ )| <span style="color: green"> done </span>| [ Climbing Stairs ]( dp/climbing-stairs.go )| |
|[121]( https://leetcode.com/problems/best-time-to-buy-and-sell-stock/ )| <span style="color: green"> done </span>| [ Best Time to Buy and Sell Stock ]( dp/best-time-to-buy-and-sell-stock.go )| |
|[53]( https://leetcode.com/problems/maximum-subarray/ )| <span style="color: green"> done </span>| [ Maximum Subarray ]( dp/maximum-subarray.go )| |
|[198]( https://leetcode.com/problems/house-robber/ )| <span style="color: green"> done </span>| [ House Robber ]( dp/house-robber.go )| |

## Design 
| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[155]( https://leetcode.com/problems/min-stack/ )| <span style="color: green"> done </span>| [ Min Stack ]( design/min-stack.go )| interfaces |
|[384]( https://leetcode.com/problems/shuffle-an-array/ )| <span style="color: green"> done </span>| [ Shuffle an Array ]( design/shuffle-an-array.go )| no need to create new arrays, can just use the same input slice |

## Math 
| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[412]( https://leetcode.com/problems/fizz-buzz/ )| <span style="color: orange "> redo </span>| [ Fizz Buzz ]( math/fizz-buzz.go )| try an optimized version with maps 
|[204]( https://leetcode.com/problems/count-primes/ )| <span style="color: orange "> </span>| [ Fizz Buzz ]( math/count-primes.go )| |

## Bits
| LC | Status | Program | Notes |
|---------|--------|---------|---------|
|[191]( https://leetcode.com/problems/number-of-1-bits/ )| <span style="color: green"> done </span>| [ Number of 1 Bits ]( bits/number-of-1-bits.go )| Learnt about Binary literals, Unsetting last set bit using x &= (x-1)|
|[461]( https://leetcode.com/problems/hamming-distance/ )| <span style="color: green"> done </span>| [ Hamming Distance ]( bits/hamming-distance.go )| interesting soln with xor |
|[190]( https://leetcode.com/problems/reverse-bits/ )| <span style="color: green"> done </span>| [ Reverse Bits ]( bits/reverse-bits.go )| interesting bit shifting/swapping logic| 

## Others 
| LC | Status | Program | Notes |
|---------|--------|---------|---------|


> template for this file

|[14](
https://leetcode.com/problems/longest-common-prefix/
)|
<span style="color: green">
done
</span>|
[
Longest Common Prefix
](
strings/longest-common-prefix.go
)|
| 