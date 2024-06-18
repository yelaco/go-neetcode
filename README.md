# Neetcode

## Note

### Bucket sort

Bucket sort is a sorting technique that involves dividing elements into various groups, or buckets. These buckets are formed by uniformly distributing the elements. Once the elements are divided into buckets, they can be sorted using any other sorting algorithm. Finally, the sorted elements are gathered together in an ordered fashion.

```python
buckets = [[] for _ in range(n)]
```

==> Elements are mapped to ```buckets``` based on some criteria, which determine the indexes.
