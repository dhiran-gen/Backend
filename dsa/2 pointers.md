Pattern
Problem Signs
Pointer Directions
Typical Use
Start-End Pointers
Data is sorted, and you’re looking for a pair, sum, or max/min
One starts at the beginning, the other at the end
Searching for a target sum, max area
Slow-Fast Pointers
You need to modify in-place or remove duplicates
One moves slowly, the other scans ahead
In-place operations, removing duplicates
Sliding Window
You are looking for substrings/subarrays that meet a condition (like length or sum)
Start and end of window move as needed
Longest substring, smallest window
Cycle Detection
Problem involves a linked list and asks about cycles
One pointer moves faster
Detecting cycles
Merge Pointers
You’re merging or comparing two sorted arrays/lists
Two pointers on separate inputs
Merging, intersection, median of arrays

Here’s a list of 15 hand-picked Two Pointer problems, divided across Easy, Medium, and Hard, categorized into the 5 core patterns you mentioned.

⸻

✅ Pattern 1: Opposite Ends (Sorted Array / Pair Search)

Use when the array is sorted, and you want to find pairs, triplets, or max/min combinations.

Easy:
	1.	Two Sum II – Input Array is Sorted (#167)
	2.	Valid Palindrome (#125)
	3.	Reverse String (#344)
	4.	Squares of a Sorted Array (#977)
	5.	Merge Two Sorted Arrays (#88)

Medium:
	1.	3Sum (#15)
	2.	Container With Most Water (#11)
	3.	Remove Duplicates from Sorted Array II (#80)
	4.	Partition Labels (#763)
	5.	Max Number of K-Sum Pairs (#1679)

Hard:
	1.	4Sum (#18)
	2.	Palindrome Pairs (#336)
	3.	Maximize Distance to Closest Person (#849)
	4.	Shortest Subarray with Sum at Least K (#862) (combines deque + 2 pointers)
	5.	Count of Range Sum (#327) (Two pointers + prefix sum)

⸻

✅ Pattern 2: Same Direction (Slow and Fast Pointers)

Used for modification in-place, skipping duplicates, etc.

Easy:
	1.	Remove Duplicates from Sorted Array (#26)
	2.	Move Zeroes (#283)
	3.	Remove Element (#27)
	4.	Middle of the Linked List (#876)
	5.	Linked List Cycle (#141)

Medium:
	1.	Linked List Cycle II (#142)
	2.	Sort Colors (#75) (Dutch National Flag)
	3.	Remove Duplicates from Sorted List II (#82)
	4.	Minimum Size Subarray Sum (#209)
	5.	Odd Even Linked List (#328)

Hard:
	1.	Reorder List (#143)
	2.	Palindrome Linked List (#234)
	3.	Reverse Nodes in k-Group (#25)
	4.	Longest Duplicate Substring (#1044) (two pointers + binary search + rolling hash)
	5.	Find the Duplicate Number (#287) (Floyd’s cycle detection)

⸻

✅ Pattern 3: Sliding Window (Dynamic Size Window)

Perfect for subarrays/substrings with constraints.

Easy:
	1.	Maximum Average Subarray I (#643)
	2.	Best Time to Buy and Sell Stock (#121)
	3.	Find Max Average Subarray II (#644)
	4.	Check If a Word Occurs As a Prefix (#1455)
	5.	Subarray Sum Equals K (Brute version)

Medium:
	1.	Longest Substring Without Repeating Characters (#3)
	2.	Permutation in String (#567)
	3.	Minimum Window Substring (#76)
	4.	Longest Repeating Character Replacement (#424)
	5.	Sliding Window Maximum (#239)

Hard:
	1.	Substring with Concatenation of All Words (#30)
	2.	Maximum Number of Vowels in a Substring (#1456)
	3.	Count Number of Nice Subarrays (#1248)
	4.	Minimum Operations to Reduce X to Zero (#1658)
	5.	Max Consecutive Ones III (#1004)

⸻

✅ Pattern 4: Fast & Slow (Cycle Detection)

Perfect for finding loops, middle elements, etc.

Easy:
	1.	Middle of the Linked List (#876)
	2.	Linked List Cycle (#141)
	3.	Happy Number (#202)
	4.	Reverse Linked List (#206)
	5.	Detect Capital (#520)

Medium:
	1.	Linked List Cycle II (#142)
	2.	Find the Duplicate Number (#287)
	3.	Longest Palindromic Substring (#5)
	4.	Reorder List (#143)
	5.	Intersection of Two Linked Lists (#160)

Hard:
	1.	Palindrome Linked List (#234)
	2.	Reverse Nodes in k-Group (#25)
	3.	Copy List with Random Pointer (#138)
	4.	Linked List Random Node (#382)
	5.	Sort List (#148)

⸻

✅ Pattern 5: Merge Two Sorted (Merge Walk)

Use when you have two sorted inputs and want to combine or compare.

Easy:
	1.	Merge Sorted Array (#88)
	2.	Intersection of Two Arrays II (#350)
	3.	Backspace String Compare (#844)
	4.	Array Partition I (#561)
	5.	Intersection of Two Arrays (#349)

Medium:
	1.	Merge Intervals (#56)
	2.	Interval List Intersections (#986)
	3.	Median of Two Sorted Arrays (#4)
	4.	Minimum Common Value (#2540)
	5.	Smallest Range Covering Elements from K Lists (#632)

Hard:
	1.	Median of Two Sorted Arrays (#4)
	2.	Smallest Range from K Lists (#632)
	3.	Count of Smaller Numbers After Self (#315)
	4.	Interval List Intersections (#986)
	5.	Minimum Window Subsequence (#727)

⸻
