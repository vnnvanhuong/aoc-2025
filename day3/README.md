# Day 3: Lobby

**Problem**
- 987654321111111 -> 98
- 811111111111119 -> 89
- 234234234234278 -> 78
- 818181911112111 -> 92

result = 357

find the max number with two digits in the given string

**Constraints**
- possitive number: 1 - 9
- could not change the order of digits. the number must be formed from left to right.
- string with 101 len

**Test cases**

input.txt

**Logical solution**

The naive approach is to check every pair of number from left to right and find the maximum of each string
two pointer
- p1: 0 ... len(s)
- p2: p1 + 1 ... len(s)

Time complexity: O(N^2). Space complexity: O(1)

```go
func totalMaxjoltage(banks []string) int {
    total := 0

    for _, b := range banks {
        n := len(b)
        m := 0
        for p1 := 0; p1 < n; p1++ {
            for p2 := p1+1; p2 < n; p2++ {
                j := strconv.Atoi(b[p1]+b[p2])
                m = max(m, j)
            }
        }

        total += m
    } 

    return total
}
```

**Optimize?**
Intuition
- Hashmap? 1 - 9
- Shrink the array with two pointer?

