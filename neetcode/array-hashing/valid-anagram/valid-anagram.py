

from itertools import count


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        countS, countT = {}, {}

        for i in range(len(s)):
            countS[s[i]] = 1 + countS.get(s[i], 0)   
            countT[s[i]] = 1 + countT.get(s[i], 0)   

        for char in countS:
            if countS[char] != countT.get(char, 0):
                return False
        
        return True