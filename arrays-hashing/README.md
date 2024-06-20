# Solutions for the Array & Hashing section

## Easy problems

## Medium problems

### Product of Array Except Self

**The idea to solve the problem is:** *the product discluding an element can be divided into two parts: left side product and right side product*

![i2](/images/poads_2.jpg)

To find the ```left product``` of the ```i-th``` element, we multiply the ```(i-1)-th``` element with its left product, starting from the ```0-th``` element with left product of 1 (first element doesn't have left side).

![i1](/images/poads_1.jpg)

The same goes for finding ```right product``` of ```i-th``` element, we multiply the ```(i+1)-th``` element with its right product.

After that, we just loop through the products and multiply them together
```i-th product except self``` = ```i-th left product``` * ```i-th right product```
