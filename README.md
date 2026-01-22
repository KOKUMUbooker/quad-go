# Quadrangle (Quad) Project

## Overview

The **Quadrangle** project implements a set of Go functions that print rectangles (quads) of a given width (`x`) and height (`y`) using specific ASCII patterns.

Each function (`QuadA` to `QuadE`) draws a rectangle with different border characters while following the same rules:

* If `x > 0` and `y > 0`, the rectangle is printed.
* If `x <= 0` or `y <= 0`, nothing is printed.
* Output is printed directly to standard output.
* Each rectangle ends with a newline (`\n`) at the end of every line.

---

## Project Structure

```
quad/
├── go.mod
├── quadA.go
├── quadB.go
├── quadC.go
├── quadD.go
├── quadE.go
└── test
    └── main.go
```

### Files Description

* **quadA.go** – Implementation of `QuadA`
* **quadB.go** – Implementation of `QuadB`
* **quadC.go** – Implementation of `QuadC`
* **quadD.go** – Implementation of `QuadD`
* **quadE.go** – Implementation of `QuadE`
* **test/main.go** – Sample test program to run and verify outputs
* **go.mod** – Go module definition

---

## Function Specifications

Each function follows the same signature:

```go
func QuadX(x, y int)
```

Where `QuadX` can be `QuadA`, `QuadB`, `QuadC`, `QuadD`, or `QuadE`.

---

## QuadA

### Pattern

* Corners: `o`
* Horizontal edges: `-`
* Vertical edges: `|`

### Example

```go
QuadA(5, 3)
```

Output:

```
o---o
|   |
o---o
```

---

## QuadB

### Pattern

* Top-left corner: `/`
* Top-right corner: `\`
* Bottom-left corner: `\`
* Bottom-right corner: `/`
* Horizontal edges: `*`
* Vertical edges: `*`

### Example

```go
QuadB(5, 3)
```

Output:

```
/***\
*   *
\***/
```

---

## QuadC

### Pattern

* Top-left corner: `A`
* Top-right corner: `A`
* Bottom-left corner: `C`
* Bottom-right corner: `C`
* Horizontal edges: `B`
* Vertical edges: `B`

### Example

```go
QuadC(5, 3)
```

Output:

```
ABBBA
B   B
CBBBC
```

---

## QuadD

### Pattern

* Top-left corner: `A`
* Top-right corner: `C`
* Bottom-left corner: `A`
* Bottom-right corner: `C`
* Horizontal edges: `B`
* Vertical edges: `B`

### Example

```go
QuadD(5, 3)
```

Output:

```
ABBBC
B   B
ABBBC
```

---

## QuadE

### Pattern

* Top-left corner: `A`
* Top-right corner: `C`
* Bottom-left corner: `C`
* Bottom-right corner: `A`
* Horizontal edges: `B`
* Vertical edges: `B`

### Example

```go
QuadE(5, 3)
```

Output:

```
ABBBC
B   B
CBBBA
```

---

## Running the Project

1. Navigate to the project root:

   ```bash
   cd quad
   ```

2. Run the test program:

   ```bash
   go run ./test
   ```

3. Modify `test/main.go` to call different `QuadX` functions with different values of `x` and `y`.

---

## Notes

* All functions are pure output functions and return no value.
* No output is produced for invalid dimensions (`x <= 0` or `y <= 0`).
* The project follows standard Go formatting and conventions.

