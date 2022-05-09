# 121. Best Time to Buy and Sell Stock [Easy]

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

 

Constraints:

    1 <= prices.length <= 105
    0 <= prices[i] <= 104

## SOLUTION 1

we use two pointers, L on day to buy, R on day to sell (in order to maximize profit, R-L).
we dont want negative profit i.e (neg R-L).

if R-L is negative that means there's gonna be a loss..so we move L to the R day and continue to move R for a positive R-L.
we compare R and L again(R-L).

if we get a positive profit (R-L= +), we update MAX_PROFIT to profit.
**if we just got a positive profit, we don't move the day we buy(L stays same), we only move day to buy(R)** so that we can reach the day we can get another profit(which might be possibly higher that our current one, and if for any reason that day doesn't exists, well we know that we have actually reached max'ed profit with our current MAX_PROFIT). **so even if our subsequent R-L is positive but less than current MAX-PROFIT, that would be considered loss too(so we still dont move L)** .(we're only **in search of a day for "MORE PROFIT"** if possible)

**WHICH MEANS THAT, L will not be moved until we get a -negative R-L,  a potential new day that  is lesser than the current L** i.e if we come across a **day that is lower than our current buy date(L), we immediately jumped there** (because there's then possibility of buying even more lower, and can rest assured that our MAX_PROFIT is still intact for now just incase we don't get any further sell day R that will give us positive R-L greater than the current MAX_PROFIT, so win-win :) )

## TAKEAWAYS
L is only moved when we get - negative R-L( L is move to the R, because we want L to be at the minimum possible)
R will continued to be moved while R-L is positive, setting MAX-PROFIT to the highest R-L allong the way until we reach the end.


### TIME: O(n)
### SPACE: O(1)