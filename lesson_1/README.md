# Go tasks one
*first tasks for students lab*
____
## Palindrome
Given the string, check if it is a palindrome.
Example
For inputString = "aabaa", the output should be solution(inputString) = true;

For inputString = "abac", the output should be solution(inputString) = false;

For inputString = "a", the output should be solution(inputString) = true.
____
## Matrix elements sum
Given matrix, a rectangular matrix of integers, where each value represents the cost of the room, your task is to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all the values that don't appear below a 0).

Example
For
 matrix = 

```
    [[0, 1, 1, 2],
    [0, 5, 0, 0],
    [2, 0, 3, 3]]
```

the output should be solution(matrix) = 9.
____
## Reverse line Parentheses
Write a function that reverses characters in (possibly nested) parentheses in the input string.
Input strings will always be well-formed with matching ()s.
Example
For inputString = "(bar)", the output should be solution(inputString) = "rab";

For inputString = "foo(bar)baz", the output should be solution(inputString) = "foorabbaz";

For inputString = "foo(bar)baz(blim)", the output should be solution(inputString) = "foorabbazmilb";

For inputString = "foo(bar(baz))blim", the output should be solution(inputString) = "foobazrabblim".

Because "foo(bar(baz))blim" becomes "foo(barzab)blim" and then "foobazrabblim".