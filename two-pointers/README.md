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

### 3Sum

**The idea to solve the problem is:** *We reuse the idea from two sum*

![i1](/images/3sum_1.jpg)

Finding a three sum ```a + b + c = 0``` would mean to find a two sum ```b + c = -a```. This reduces our problem to two sum, in which we need to find the pairs that sum up to target ```-a```

In problem ```Two Sum II```, we use two pointers to solve, which is very efficient. We will also use two pointers technique for this problem. But it requires the array to be sorted first. We can use ```hashmap``` instead but it requires more space complexity and isn't actually better than using two pointer in reducing time complexity for this problem (we will talk about that later)

![i2](/images/3sum_2.jpg)

So we sort the array first. Then we can just loop through the array and find two sum pairs for the current element.

Note that the triplets have to be **unique**, so we will skip the duplicate elements (targets) while looping. We also have to skip the duplicate two sum pairs for each unique target. To do that, when we find a pair, we keep increasing the ```left``` pointer till the current ```nums[left]``` is different from the previous ```nums[left]```. This effectively avoids all the possible duplicate triplets for the previous ```nums[left]``` element. And since the ```nums[left]``` is different (bigger than previous), the sum will always be positive until the ```right``` index is decreased.