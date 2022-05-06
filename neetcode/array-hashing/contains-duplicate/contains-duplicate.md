Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.


## SOLUTION 1

take each element and compare with each element in the array

TIME: O(n2)
SPACE: O(1)


## SOLUTION 2

sort the array, then compare each adjacent element

TIME: O(nlogn)
SPACE: O(1)


## SOLUTION 3 [ACCEPTED]

use an hashset to recored element occcurence(extra space)
hashset allows O(1) insert & lookup
we insert an element if it doesn't exist, we then confirm it has a dup when we encouter it again(elem : elem)

TIME: O(n)
SPACE: O(n)