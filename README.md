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

### Reading integers from a stream

```go
file, err := os.Open("numbers.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

stats, err := summary.StreamInt(file)
if err != nil {
    log.Fatal(err)
}
```

Example input:

```text
10
20
30
40
```

### Reading floating-point numbers

```go
file, err := os.Open("values.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

stats, err := summary.StreamFloat(file, 64)
if err != nil {
    log.Fatal(err)
}
```

## API

### Slice functions

```go
summary.Ints(...)
summary.Floats(...)
```

### Stream functions

```go
summary.StreamInts(...)
summary.StreamFloats(...)
```

## License

MIT License.
