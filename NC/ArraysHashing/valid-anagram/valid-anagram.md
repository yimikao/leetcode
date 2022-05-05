Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false
 

Constraints:

    1 <= s.length, t.length <= 5 * 104
    s and t consist of lowercase English letters.


## SOLUTION 1

insert every element's occurence into an hashmap S & T. Then compare them

TIME: O(S+T) i.e O(n)
SPACE: O(S+T) i.e O(n)


how would you then solve it without additional memory i.e SPACE: O(1)
- we sort both string and compare them ==
- leaving us with: 
    TIME: O(nlogn) if we use a fast sort algo
    SPACE: O(1) most sorting algo use constant space & interviers assume sorting takes no extra space