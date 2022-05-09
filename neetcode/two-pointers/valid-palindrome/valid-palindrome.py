from typing import List


class Solution:
    def isPalindrome(self, s: str) -> bool:
        newStr = ""

        for c in s:
            if s.isalnum():
                newStr += c.lower()
        return newStr == newStr[::-1] #new string and it's reverse are same?


class Solution2: #using pointers
    def isPalindrome(self, s: str) -> bool:
