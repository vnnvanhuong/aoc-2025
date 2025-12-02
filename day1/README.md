# Day 1

## Understand the problem
p = 50
rotations = [L68, L30, R48, L5, R60, L55, L1, L99, R14, L82] => pwd = 3

explanation:
- L68 -> 50 - 68 = -18, 100 - 18 = 82
- L30 -> 82-30 = 52
- R48 -> 52 + 48 = 100, pwd=1
- L5  -> 100 - 5 = 95
- R60 -> 95 + 60 = 155, 155 - 100 = 55
- L55 -> 55 - 55 = 0, pwd=2
- L1  -> 0 - 1 = -1, 100 - 1 = 99
- L99 -> 99 - 99 = 0, pwd=3
- R14 -> 0 + 14 = 14
- L82 -> 14 - 82 = -68, 100 - 68 = 32

## Constraints
- turns is only from 1 - 99
- only moving left or right
- there is always an answer

## Test cases
- p = 50
rotations = [L68, L30, R48, L5, R60, L55, L1, L99, R14, L82] => pwd = 3

- p=50
rotations = [L68, L30, R48, L5, R60, L55] => pwd = 2

- p=50
rotations = [L50] => pwd = 1

- p=50
rotations = [R50] => pwd = 1

## Logical solution

find the formular:

- if moving left (L), new position = current position - turns, if turns < 0, new position = 100 - new position
- if moving right (R), new position = current position + turns, if turns > 100, new position = new position - 100
- if new position = 100 or new position = 0, pwd++
- if new position = 100, new position = 0

## Write the code

```go
// iterate through the rotation items
// parse the direction and turns
func pwd(p int, rotations []string) int {
    pwd := 0
    for _, r := range rotations {
        direction, turns := parse(r)
        if direction = "L" {
            p = p - turns
            if p < 0 {
                p = p + 100
            }
        } else {
            p = p + turns
            if p > 100 {
                p = p - 100
            }
        }

        if p == 100 || p == 0 {
            pwd++
        }

        if p == 100 {
            p == 0
        }

    }

    return pwd
}

func parse(rotation string) (string, int) {
    directions := rotation[0]
    turns, _ := strconv.Atoi(rotation[1:])

    return direction, turns
}

```

## Test the code (dry-run)

[L68, L30, R48, L5, R60, L55]

```go
func pwd(rotations []string) int {
    p := 50
    pwd := 0
    for _, r := range rotations {
        direction, turns := parse(r) // L, 55
        if direction = "L" {
            p = p - turns // p = 55 - 55 = 0
            if p < 0 {
                p = p + 100 // p = -5 + 100 = 95
            }
        } else {
            p = p + turns // -5 + 60 = 55
            if p > 100 {
                p = p - 100
            }
        }

        if p == 100 || p == 0 {
            pwd++ // 2
        }

        if p == 100 {
            p == 0 // 0
        }

    }

    return pwd // 2
}

func parse(rotation string) (string, int) {
    directions := rotation[0]
    turns, _ := strconv.Atoi(rotation[1:])

    return direction, turns
}

```
