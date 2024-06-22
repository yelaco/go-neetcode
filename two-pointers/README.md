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

The same goes for finding ```right product``` of ```i-th``` element, we multiply the ```(i+1)-th``` element with its right product.

After that, we just loop through the products and multiply them together
```i-th product except self``` = ```i-th left product``` * ```i-th right product```
