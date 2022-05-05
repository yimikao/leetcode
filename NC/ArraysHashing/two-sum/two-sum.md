Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## EXAMPLES 
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]

## SOLUTION 1
we pick each element and check down the array comparing each elements to what we pick.this is bruteforce

TIME: O(n2)
SPACE: O(1)


## SOLUTION 2
we don't add the entire array initially because we cant reuse num i.e we cant add same num twice into the array

the clever way to do it is: take [2,1,5,3] target 4

- we say how hashmap is initially empty
e.g starting with 2, we do 4-2= 2. since it doesnt already exists, we add it into our HM
-> 2 : 0

<!-- (NOTE: ) -->
(NOTE: we will make every num lookup able in O(1) by inserting to HM
hashmaps are a cheat/quick register we can use for easy insert/lookup. think about how u can exploit that feature
)
 so:
* check if target - num (our answer maker) exists in HM 
* if not make sure num isn't in HM already before inserting (we are checking if element already exists before adding to avoid duplicates)

(**SO WE ARE CHECKING POSSIBILITY OF OUR ANSWER, LOOKUP AND INSERT AT ONE ITERATION..AND OFCOURSE THE INDEX, THAT'S PRETTY COOL**)

4-1 = 3, 3(answer) doesn't exist in HM yet(i.e we can confirm that we have an answer) so we add 1 to our HM
-> 1 : 1

4-5= -1
-> 5:1

4-3= 1, 1exists in our HM,
-> 3:3, return [index 1 and 3]

*(we are always guaranteed that no matter the num we are on, if it has a sum we would have added it into the HM since we've had reached it)*


- **TIME: O(n)**
- **SPACE: O(n)**