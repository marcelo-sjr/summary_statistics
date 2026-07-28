[![Go Reference](https://pkg.go.dev/badge/github.com/marcelo-sjr/summary_statistics.svg)](https://pkg.go.dev/github.com/marcelo-sjr/summary_statistics)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/marcelo-sjr/summary_statistics)](go.mod)
# Summary

`summary` is a lightweight Go library for computing summary statistics from numeric collections and streams.

It provides an API inspired by Java's `IntSummaryStatistics` and `DoubleSummaryStatistics`, while following idiomatic Go design.

## Features

- Generic support for integer and floating-point slices
- Read numbers directly from any `io.Reader`
- Computes:
  - Count
  - Sum
  - Average
  - Minimum
  - Maximum
- Invalid lines are ignored when reading streams
- Small, dependency-free package

## Installation

```bash
go get github.com/marcelo-sjr/summary_statistics
```

## Usage

### Example

```go
stats := summary.Int(10, 20, 30, 40)

fmt.Println(stats.Count)
fmt.Println(stats.Sum)
fmt.Println(stats.Avg)
fmt.Println(stats.Min)
fmt.Println(stats.Max)
```

## Reading from streams

Both `StreamInt` and `StreamFloat` accept any `io.Reader` and a
`bufio.SplitFunc`, allowing you to control how the input is tokenized.

### One value per line

```go
file, err := os.Open("numbers.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

stats, err := summary.StreamInt(file, bufio.ScanLines)
```

Input:

```text
10
20
30
40
```

---

### Space-separated values

```go
file, err := os.Open("numbers.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

stats, err := summary.StreamInt(file, bufio.ScanWords)
```

Input:

```text
10 20 30 40
```

---

### Mixed input

Using `bufio.ScanWords`, invalid tokens are ignored.

Input:

```text
10
20
hello
30
40 foo
50
```

Only the numeric tokens are included in the statistics.

## API

### Slice functions

```go
summary.Ints[T integer](t ...T) IntSummary
summary.Floats[T float](t ...T) FloatSummary
```

### Stream functions

```go
summary.StreamInts(r io.Reader, split bufio.SplitFunc) (IntSummary, error)
summary.StreamFloats(r io.Reader, bitSize int, split bufio.SplitFunc) (FloatSummary, error)
```

## License

MIT License.
