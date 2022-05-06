

"""
    Created by damilolaolayinka on 4/4/22.

    - Arrays is a contigous block of memory
    - Reviewing & Updating Arrays takes O(1) time
    - Insertion & Deletion take O(n) time : which both require allocation of new array

"""


#********* PATTERN *****#
# PARTITIONING ARRAYS FOR FASTER SORTING [ O(1) not O(n) ]
# SWAPPING ELEMENTS IN PARTITIONS (I.E SUBARRAYS)

"""
    Exercise 1:
        Reorder array entries so that the even entries appear first.
        you are required to solve it withoutallocating additional storage.
"""

def even_odd(A):
   
    next_even, next_odd = 0, len(A) - 1  # partition array

    while next_even < next_odd:
        if next_even % 2 == 0:
            next_even +=1  # move forward
        else:
            A[next_even], A[next_odd] = A[next_odd], A[next_even] #swap values


'''
    Ex2: Dutch National Flag Problem
        Write a program that takes an array A and an index i into A, and rearranges the elements such
        that all elements less than A[i] (the "pivot") appear first, followed by elements equal to the pivot,
        followed by elements greater than the pivot.

        Hint: Think about the partition step in quicksort.

        Bruteforce: The problem is trivial to solve with O(n) additional space/ where n is the length of A.
        We form three lists, namely, elements less than the pivot, elements equal to the pivot, and elements
        greater than the pivot. Consequently, we write these values into A. The time complexity is O(n).

'''