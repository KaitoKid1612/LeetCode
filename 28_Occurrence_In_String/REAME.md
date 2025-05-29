# 28. Find the Index of the First Occurrence in a String

## Problem Statement
Given two strings `haystack` and `needle`
Return the index of the first occurrence of `needle` in `haystack`, or -1 if `needle` is not part of `haystack`.

## Example
```Example 1:
> Input: haystack = "hello", needle = "ll"
> Output: 2
> Explanation: The substring "ll" occurs at index 2 in the string "hello".
```

```Example 2:
> Input: haystack = "aaaaa", needle = "bba"
> Output: -1
> Explanation: The substring "bba" is not found in the string "aaaaa", so the output is -1.
```

```Example 3:
> Input: haystack = "", needle = ""
> Output: 0
> Explanation: An empty string is considered to be found at index 0 in another empty string.
```

```Example 4:
> Input: haystack = "sadbutsad", needle = "sad"
> Output: 0
> Explanation: The substring "sad" occurs at index 0 in the string "sadbutsad".
```
