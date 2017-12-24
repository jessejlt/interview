About
---

A collection of solutions to questions asked during the coding portion of an interview. Just trying to keep this particular set of skills fresh.

Questions
---

`partial_products.go`

> Asked by Google

Given a sorted set of integers [1, 2, 3, 4], produce a set of partial products such that each index contains the product of the set, sans the current index.

**Restrictions**
1. You cannot use division
2. Solution must be O(n)

**Example**
Input:  [1, 2, 3, 4]
Output: [24, 12, 8, 6]

`topo_hydro_volume.go`

> Asked by Google

You're given a list of integers that represent the topology of an island; find the total volume of water your island can hold. 

**Example**
Input: [1, 3, 5, 1, 7, 2, 3, 6]
Output: 11

```
     +          +--+
    7|          |  |     +--+
     |          |  |     |  |
    6|          |  |     |  |
     |          |  |     |  |
    5|    +--+  |  |     |  |
     |    |  |  |  |     |  |
    4|    |  |  |  |     |  |
     |    |  |  |  |  +--+  |
    3| +--+  |  |  |  |  |  |
     | |  |  |  |  |  |  |  |
    2| |  |  |  |  +--+  |  |
     | |  |  |  |  |  |  |  |
    1+-+  |  +--+  |  |  |  |
     | |  |  |  |  |  |  |  |
    0| |  |  |  |  |  |  |  |
+------|--|--|--|--|--|--|--|--+
     |0  1  2  3  4  5  6  7
     +
```
