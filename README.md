[![Go Reference](https://pkg.go.dev/badge/github.com/marcelo-sjr/summary_statistics.svg)](https://pkg.go.dev/github.com/marcelo-sjr/summary_statistics)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/marcelo-sjr/summary_statistics)](go.mod)

# Summary

`summary` is a lightweight Go library for computing summary statistics from numeric collections and streams.

It provides an API inspired by Java's `IntSummaryStatistics` and `DoubleSummaryStatistics`, while following idiomatic Go design.

## Features

- Generic support for integer and floating-point slices
- Read numbers directly from any `io.Reader`
- Configurable tokenization using `bufio.SplitFunc`
- `bufio.ScanWords` is used by default when no split function is provided
- Computes:
  - Count
  - Sum
  - Average
  - Minimum
  - Maximum
- Invalid tokens are ignored when reading streams
- Dependency-free

## Installation

```bash
go get github.com/marcelo-sjr/summary_statistics
```

## Usage

### Slice

```go
stats := summary.Ints(10, 20, 30, 40)

fmt.Println(stats.Count)
fmt.Println(stats.Sum)
fmt.Println(stats.Avg)
fmt.Println(stats.Min)
fmt.Println(stats.Max)
```

### Floating-point slice

```go
stats := summary.Floats(1.5, 2.5, 3.5)

fmt.Println(stats.Count)
fmt.Println(stats.Sum)
fmt.Println(stats.Avg)
fmt.Println(stats.Min)
fmt.Println(stats.Max)
```

## Reading from streams

`ReadInts` and `ReadFloats` accept any `io.Reader`.

By default, input is tokenized using `bufio.ScanWords`. You may provide a custom `bufio.SplitFunc` when different tokenization is required.

### Default behavior

```go
file, err := os.Open("numbers.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

stats, err := summary.ReadInts(file, nil)
if err != nil {
	log.Fatal(err)
}
```

Input:

```text
10 20 30 40
```

---

### One value per line

```go
stats, err := summary.ReadInts(file, bufio.ScanLines)
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
stats, err := summary.ReadInts(file, bufio.ScanWords)
```

Input:

```text
10 20 30 40
```

---

### Mixed input

Invalid tokens are ignored.

```text
10
20
hello
30
40 foo
50
```

Only numeric tokens are included in the statistics.

## API

### Slice functions

```go
summary.Ints[T integer](values ...T) IntSummary
summary.Floats[T float](values ...T) FloatSummary
```

### Reader functions

```go
summary.ReadInts(r io.Reader, split bufio.SplitFunc) (IntSummary, error)
summary.ReadFloats(r io.Reader, split bufio.SplitFunc) (FloatSummary, error)
```

## License

MIT License.
