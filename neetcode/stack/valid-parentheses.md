# 20. Valid Parentheses [Easy]

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

 

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

 

Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.



# SOLUTION
**when we come across an opening bracket in the string, we add it to the stack** for comparison and we can continue to add if it is opening we're getting, because (((())))) is valid.
if not we first **check that stack is not empty** because we have to make sure that we already have an opening one that we wanna comepare against **use the closing to lookup it's opening counterpart to check if it matches the opening in the stack**, and if it does, we remove that latest opening that we just added.


## THINGS TO USE
1) we use a stack to keep track of /single out current opening bracket..easy to peek and pop at the end of a stack.

2) we use a **hashmap** to check: 
* if a bracket is opening or closing
* check if opening bracket in stack matches next bracket(which is expected to be closing)


## HOW IT WORKS

## EXAMPLE 1: "{()}("

* opening
* stack => ["{"]
* opening
* stack => ["{", "("]
* closing
* check stack != [] and  ")":"(" == stack(-1), which is True, continue (NOTE THAT a "}" or "?"  will be caught here and False will be returned because stack will be empty or stack[-1] i.e openingBracket will not equal closeToOpen("currentClosingBracket") respectively)

* stack => ["{"]
* check "}":"{" == stack(-1), which is True, continue
* opening
* stack => ["("]

* which is False, because we have an opharn opening i.e stack != [] (i.e **"not stack"** will give us false), we could have gotten True if stack == [] because **not stack** will give us True


## EXAMPLE 2: "{()})"

the last closing bracket will be at:
    * if stack and stack[-1]......
because stack will be empty and condition will be if (stack) will be if (false)


# TIME : O(n)
# SPACE: O(n) stack could go to end
