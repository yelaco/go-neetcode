# Solutions for the Two Pointers section

## Easy problems

## Medium problems

### Two Sum II

**The idea to solve the problem is:** *We use 2 pointers to keep track of the current sum of 2 elements bounding the 2 elements we need to find*

![i2](/images/two_sum_II_1.jpg)

Starting with a pointer ```i``` to first element and pointer ```j``` to last element of the array, we have

```python
sum = nums[i] + nums[j]
```

![i1](/images/two_sum_II_2.jpg)

If the sum is larger, we know that we have to decrease ```j``` because the elements after ```j-th``` are all bigger than ```j-th``` element, which's not gonna make the ```sum``` any closer to the target.

Similarly, if the sum is smaller, we know that we have to increase ```i``` because the elements before ```i-th``` are all bigger than ```i-th``` element, which's not gonna make the ```sum``` any closer to the target.

Because it is guaranteed to have only 1 solution, once the left pointer ```i``` reaches a solution element, the sum will always bigger than the ```target``` until the right pointer ```j``` also reaches the other solution element, and vice versa. At this point, we return the solution.
