# 2894. Divisible and Non-divisible Sums Difference

## Problem Statement
You are given positive integers `n` and `m`.
Define two integers as follows:

- `num1`: The sum of all integers in the range ```[1, n]``` (both inclusive) that are **not divisible** by `m`.
- `num2`: The sum of all integers in the range ```[1, n]``` (both inclusive) that are **divisible** by `m`.
Return the integer `num1 - num2`.
## Example
```Example 1:
> Input: n = 10, m = 3
> Output: 19
> Explanation: In the given example:
> - The integers from 1 to 10 that are not divisible by 3 are: 1, 2, 4, 5, 7, 8, 10. num1 = 1 + 2 + 4 + 5 + 7 + 8 + 10 = 37.
> - The integers from 1 to 10 that are divisible by 3 are: 3, 6, 9. num2 = 3 + 6 + 9 = 18.
> - The difference is num1 - num2 = 37 - 18 = 19.
```

```Example 2:
> Input: n = 5, m = 6
> Output: 15
> Explanation: In the given example:
> - The integers from 1 to 5 that are not divisible by 6 are: 1, 2, 3, 4, 5. num1 = 1 + 2 + 3 + 4 + 5 = 15.
> - The integers from 1 to 5 that are divisible by 6 are: none. num2 = 0.
> - The difference is num1 - num2 = 15 - 0 = 15.
```

```Example 3:
> Input: n = 5, m = 1
> Output: -15
> Explanation: In the given example:
> - The integers from 1 to 5 that are not divisible by 1 are: none. num1 = 0.
> - The integers from 1 to 5 that are divisible by 1 are: 1, 2, 3, 4, 5. num2 = 1 + 2 + 3 + 4 + 5 = 15.
> - The difference is num1 - num2 = 0 - 15 = -15.
```