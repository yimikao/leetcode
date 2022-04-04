

"""
    Created by damilolaolayinka on 4/4/22.

    - Arrays is a contigous block of memory
    - Reviewing & Updating Arrays takes O(1) time
    - Insertion & Deletion take O(n) time : which both require allocation of new array

"""



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
