# Day 2: Gift Shop

## Table of Contents
- [Understand the problem](#understand-the-problem)
- [Contraints](#constraints)
- [Write out test cases](#write-out-test-cases)
- [Logical solution](#logical-solution)
- [Write code](#write-code)
- [Dry run](#dry-run)
- [Time and space complexity](#calculate-time-and-space-complexity)
- [Optimize](#optimize)

## Understand the problem
- 11-22 -> [11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 21] -> 11, 22
- 95-115 -> [95, 96, ..., 99, 100, ..., 115] -> 99
- 998-1012 -> 1010
- 1188511880-1188511890 -> 1188511885
- 222220-222224 -> 22222
- 1698522-1698528 -> n/a
- 446443-446449 -> 446446
- 38593856-38593862 -? 38593859
- 565653-565659 -> n/a
- 824824821-824824827 -> n/a
- 2121212118-2121212124 -> n/a


## Constraints
- None of the numbers have leading zeroes.
- integer? no, maxinum number of the input 9 milions --> float


## Write out test cases
- 11-22 -> 33
- 11-22,95-115,998-1012 -> 11 + 12 + 99 + 1010 = 1142

## Logical solution
A naive approach is to iterate each range and find an invalid ID
What is an invalid ID? Any ID which is only of some sequence of digits repeated twice.
- If len of a number is even --> no
- If len of a number is odd, check if left half equals to right half

## Write code
```go
package day2

type Range struct {
    Min int
    Max int
}

func TotalOfInvalidIDs(ranges []Range) int {
    total := 0

    // 11-22
    for _, r := range ranges {
        for i := r.Min; i <= r.Max; i++ { 
            s, _ := strconv.Aiot(i) // 13
            if len(s) % 2 != 0 {
                continue
            }

            l := s[:len(s)/2] // 1
            r := s[len(s)/2+1: len(s)] // 2
            if l != r {
                continue
            }

            total += i // 1
        }
    }

    return total
}
```
